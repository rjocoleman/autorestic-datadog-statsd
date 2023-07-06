package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/alecthomas/kong"
)

var (
	version = "dev"     // -ldflags "-X main.version=1.0.0"
	commit  = "none"    // -ldflags "-X main.commit=b54296e"
	date    = "unknown" // -ldflags "-X main.date=$(date)"
)

type CLI struct {
	Version kong.VersionFlag

	State            string `kong:"required,help='State of the backup process.',enum='before,after,success,failure'"`
	EventMessage     string `kong:"optional,name='event-message',help='Additional message to add to the event.'"`
	SendMetrics      bool   `kong:"optional,name='send-metrics',default='false',help='Flag to send metrics.'"`
	SendServiceCheck bool   `kong:"optional,name='send-service-check',default='false',help='Flag to send service check.'"`
	StatsdAddress    string `kong:"optional,name='statsd-address',env='DD_DOGSTATSD_URL',help='StatsD Address.'"`
	Location         string `kong:"optional,name='location',env='AUTORESTIC_LOCATION',help='Location for the backup.'"`
}

type StatsDClient interface {
	Event(*statsd.Event) error
	Gauge(name string, value float64, tags []string, rate float64) error
	ServiceCheck(*statsd.ServiceCheck) error
}

var cli CLI

func main() {
	kong.Parse(&cli, kong.Vars{
		"version": fmt.Sprintf("autorestic-datadog-statsd: %s, commit %s, built at %s", version, commit, date),
	})

	if cli.StatsdAddress == "" {
		log.Fatal("DD_DOGSTATSD_URL environment variable or CLI parameter not set")
	}

	client, err := statsd.New(cli.StatsdAddress)
	if err != nil {
		log.Fatal(err)
	}

	if cli.Location == "" {
		log.Fatal("AUTORESTIC_LOCATION environment variable or CLI parameter not set")
	}

	backupMetrics := []string{
		"FILES_ADDED",
		"FILES_CHANGED",
		"FILES_UNMODIFIED",
		"DIRS_ADDED",
		"DIRS_CHANGED",
		"DIRS_UNMODIFIED",
		"ADDED_SIZE",
		"PROCESSED_FILES",
		"PROCESSED_SIZE",
		"PROCESSED_DURATION",
	}

	backends := getBackends(backupMetrics)

	for backend := range backends {
		lowerBackend := strings.ToLower(backend)

		tags := generateTags(cli.Location, lowerBackend)

		sendEvent(client, cli.Location, lowerBackend, tags)
		if cli.SendServiceCheck {
			sendServiceCheck(client, cli.State, cli.Location, lowerBackend, tags)
		}
		if cli.State == "success" || cli.State == "after" {
			if cli.SendMetrics {
				sendMetrics(client, backupMetrics, cli.Location, lowerBackend, tags)
			}
		}
	}
}

func getBackends(backupMetrics []string) map[string]bool {
	backends := map[string]bool{}
	prefix := "AUTORESTIC_"

	// define a regex pattern that matches a string ending with an alphanumeric substring containing at least one letter
	regexPattern := regexp.MustCompile(`^` + prefix + `([A-Z0-9_]*[A-Z][A-Z0-9]*)_([A-Z0-9]*[A-Z][A-Z0-9]*)$`)

	for _, element := range os.Environ() {
		variable := strings.SplitN(element, "=", 2)[0] // Get the environment variable name
		if strings.HasPrefix(variable, prefix) {
			matches := regexPattern.FindStringSubmatch(variable)
			if matches != nil && len(matches) > 2 {
				metricName := matches[1]
				_, isMetric := find(backupMetrics, metricName)
				if isMetric {
					backendName := matches[2]
					backends[backendName] = true
				}
			}
		}
	}

	return backends
}

func generateTags(app string, backend string) []string {
	appTag := fmt.Sprintf("app:%s", app)
	backendTag := fmt.Sprintf("backend:%s", backend)
	return []string{appTag, backendTag, "backup", "restic", "autorestic"}
}

func sendEvent(client StatsDClient, app string, backend string, tags []string) {
	eventTitle := fmt.Sprintf("Backup %s for app %s on backend %s", cli.State, app, backend)
	snapshotID := os.Getenv(fmt.Sprintf("AUTORESTIC_SNAPSHOT_ID_%s", strings.ToUpper(backend)))
	parentSnapshotID := os.Getenv(fmt.Sprintf("AUTORESTIC_PARENT_SNAPSHOT_ID_%s", strings.ToUpper(backend)))
	eventText := fmt.Sprintf("The backup state is %s. Snapshot ID: %s. Parent Snapshot ID: %s. %s", cli.State, snapshotID, parentSnapshotID, cli.EventMessage)
	tags = append(tags, fmt.Sprintf("snapshot_id:%s", snapshotID), fmt.Sprintf("parent_snapshot_id:%s", parentSnapshotID))

	event := &statsd.Event{
		Title:          eventTitle,
		Text:           eventText,
		Priority:       statsd.Normal,
		AlertType:      statsd.Info,
		AggregationKey: app,
		Tags:           tags,
	}

	switch cli.State {
	case "success":
		event.AlertType = statsd.Success
	case "failure":
		event.AlertType = statsd.Error
	}

	log.Println("Creating event", eventTitle)

	err := client.Event(event)
	if err != nil {
		log.Println("Error sending event: ", err)
	}
}

func sendServiceCheck(client StatsDClient, state, app, backend string, tags []string) {
	snapshotID := os.Getenv(fmt.Sprintf("AUTORESTIC_SNAPSHOT_ID_%s", strings.ToUpper(backend)))
	parentSnapshotID := os.Getenv(fmt.Sprintf("AUTORESTIC_PARENT_SNAPSHOT_ID_%s", strings.ToUpper(backend)))
	tags = append(tags, fmt.Sprintf("snapshot_id:%s", snapshotID), fmt.Sprintf("parent_snapshot_id:%s", parentSnapshotID))

	status := statsd.Ok
	message := fmt.Sprintf("Backup %s for app %s on backend %s", state, app, backend)

	switch state {
	case "before":
		status = statsd.Warn
	case "failure":
		status = statsd.Critical
	}

	sc := &statsd.ServiceCheck{
		Name:    "autorestic.backup",
		Status:  status,
		Tags:    tags,
		Message: message,
	}

	err := client.ServiceCheck(sc)
	if err != nil {
		log.Printf("Failed to report service check: %s", err)
	} else {
		log.Printf("Reported service check with status %d and message %s", status, message)
	}
}

func sendMetrics(client StatsDClient, backupMetrics []string, app string, backend string, tags []string) {
	for _, metric := range backupMetrics {
		valueStr := os.Getenv(fmt.Sprintf("AUTORESTIC_%s_%s", metric, strings.ToUpper(backend)))
		if valueStr != "" {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Println("Error converting metric value to int: ", err)
				continue
			}

			log.Println("Sending metric", metric, "for backend", backend, "with value", value)
			err = client.Gauge(fmt.Sprintf("restic.%s.%s.%s", app, backend, strings.ToLower(metric)), float64(value), tags, 1)
			if err != nil {
				log.Println("Error sending metric: ", err)
			}
		}
	}
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
