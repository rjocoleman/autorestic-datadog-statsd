package main

import (
	"testing"

	"github.com/DataDog/datadog-go/v5/statsd"
	mock_statsd "github.com/rjocoleman/autorestic-dd-notify/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetBackends(t *testing.T) {
	t.Setenv("AUTORESTIC_FILES_ADDED_BACKEND1", "1")
	t.Setenv("AUTORESTIC_FILES_CHANGED_BACKEND1", "2")
	t.Setenv("AUTORESTIC_FILES_UNMODIFIED_BACKEND2", "3")

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

	expectedBackends := map[string]bool{
		"BACKEND1": true,
		"BACKEND2": true,
	}

	actualBackends := getBackends(backupMetrics)

	assert.Equal(t, expectedBackends, actualBackends)
}

func TestGenerateTags(t *testing.T) {
	app := "app1"
	backend := "backend1"

	expectedTags := []string{
		"app:app1",
		"backend:backend1",
		"backup",
		"restic",
		"autorestic",
	}

	actualTags := generateTags(app, backend)

	assert.Equal(t, expectedTags, actualTags)
}

func TestSendEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_statsd.NewMockClientInterface(ctrl)

	t.Setenv("AUTORESTIC_SNAPSHOT_ID_BACKEND1", "snapshot1")
	t.Setenv("AUTORESTIC_PARENT_SNAPSHOT_ID_BACKEND1", "parentSnapshot1")

	cli = CLI{
		State:        "success",
		EventMessage: "Successful backup",
	}

	expectedEvent := &statsd.Event{
		Title:          "Backup success for app app1 on backend backend1",
		Text:           "The backup state is success. Snapshot ID: snapshot1. Parent Snapshot ID: parentSnapshot1. Successful backup",
		Priority:       statsd.Normal,
		AlertType:      statsd.Success,
		AggregationKey: "app1",
		Tags:           []string{"app:app1", "backend:backend1", "snapshot_id:snapshot1", "parent_snapshot_id:parentSnapshot1"},
	}

	client.EXPECT().Event(expectedEvent).Return(nil)

	sendEvent(client, "app1", "backend1", []string{"app:app1", "backend:backend1"})
}

func TestSendServiceCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_statsd.NewMockClientInterface(ctrl)

	t.Setenv("AUTORESTIC_SNAPSHOT_ID_BACKEND1", "snapshot1")
	t.Setenv("AUTORESTIC_PARENT_SNAPSHOT_ID_BACKEND1", "parentSnapshot1")

	cli = CLI{
		State: "before",
	}

	expectedServiceCheck := &statsd.ServiceCheck{
		Name:    "autorestic.backup",
		Status:  statsd.Warn,
		Message: "Backup before for app app1 on backend backend1",
		Tags:    []string{"app:app1", "backend:backend1", "snapshot_id:snapshot1", "parent_snapshot_id:parentSnapshot1"},
	}

	client.EXPECT().ServiceCheck(expectedServiceCheck).Return(nil)

	sendServiceCheck(client, cli.State, "app1", "backend1", []string{"app:app1", "backend:backend1"})
}

func TestSendMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_statsd.NewMockClientInterface(ctrl)

	t.Setenv("AUTORESTIC_FILES_ADDED_BACKEND1", "1")
	t.Setenv("AUTORESTIC_FILES_CHANGED_BACKEND1", "2")
	t.Setenv("AUTORESTIC_FILES_UNMODIFIED_BACKEND1", "3")

	backupMetrics := []string{
		"FILES_ADDED",
		"FILES_CHANGED",
		"FILES_UNMODIFIED",
	}

	gomock.InOrder(
		client.EXPECT().Gauge(gomock.Any(), float64(1), gomock.Any(), gomock.Any()),
		client.EXPECT().Gauge(gomock.Any(), float64(2), gomock.Any(), gomock.Any()),
		client.EXPECT().Gauge(gomock.Any(), float64(3), gomock.Any(), gomock.Any()),
	)

	sendMetrics(client, backupMetrics, "app1", "backend1", []string{"app:app1", "backend:backend1"})
}

func TestFind(t *testing.T) {
	elements := []string{"el1", "el2", "el3"}

	index, found := find(elements, "el2")

	assert.True(t, found)
	assert.Equal(t, 1, index)

	index, found = find(elements, "el4")

	assert.False(t, found)
	assert.Equal(t, -1, index)
}
