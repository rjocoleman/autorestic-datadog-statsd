# Autorestic Datadog Notifier

This CLI application (`autorestic-dd-notify`) is designed to work with [Autorestic](https://github.com/cupcakearmy/autorestic) and [Datadog's StatsD](https://docs.datadoghq.com/developers/dogstatsd/) to report metrics for backup events.

It enables sending custom events and metrics for different backup states (before, after, success, failure) to Datadog using the StatsD protocol. This tool is particularly useful for gaining insights into your Autorestic backup processes through the visualization and alerting capabilities of Datadog.

## Prerequisites

1. [Autorestic](https://github.com/cupcakearmy/autorestic) - This application is designed to work with Autorestic, it uses the environment variables provided by Autorestic to gather metrics.

2. [Datadog Agent](https://docs.datadoghq.com/agent/) - Ensure the Datadog Agent is running either on your machine or somewhere accessible via a network.

## Usage

This application works as a command-line tool that sends data to Datadog based on the provided arguments. It's meant to be used with Autorestic's backup hooks.

The primary CLI arguments are:

- `--state` - to specify the backup state (before, after, success, failure).
- `--event-message` - to specify the message that will be attached to the Datadog event.
- `--send-event` - to specify that an event should be sent to Datadog.
- `--send-metrics` - to specify that metrics should be sent to Datadog (only valid for "after" and "success" states).
- `--send-service-check` - to specify that a service check should be sent to Datadog.

The application reads `DD_DOGSTATSD_URL` and `AUTORESTIC_LOCATION` from either environment variables or CLI parameters.

## Command Examples

Here are some examples of how you could use this tool with different backup hooks:

1. Before backup:

```shell
./autorestic-dd-notify --state before --send-event --event-message "Before backup hook running"
```

2. After backup:

```shell
./autorestic-dd-notify --state after --send-event --event-message "Finished backup process" --send-metrics
```

3. Backup failure:

```shell
./autorestic-dd-notify --state failure --send-event --event-message "Backup process failed" --send-service-check
```

4. Backup success:

```shell
./autorestic-dd-notify --state success --send-event --event-message "Backup process succeeded" --send-metrics --send-service-check
```

## Autorestic Example

To integrate with Autorestic, you might use hooks like this in your Autorestic config:

```yaml
locations:
  my-location:
    from: /data
    to: my-backend
    hooks:
      before:
        - autorestic-dd-notify --state before --event-message "Backup before hook: starting One" --send-event
        - echo "One"
        - autorestic-dd-notify --state before --event-message "Backup before hook: completed One" --send-event
        - echo "Two"
        - echo "Three"
        - autorestic-dd-notify --state before --event-message "Backup before hooks finished" --send-event
      after:
        - echo "Byte"
        - autorestic-dd-notify --state after --event-message "Finished backup process" --send-event --send-metrics
      failure:
        - echo "Something went wrong"
        - autorestic-dd-notify --state failure --event-message "Backup process failed" --send-event --send-service-check
      success:
        - echo "Well done!"
        - autorestic-dd-notify --state success --event-message "Backup process succeeded" --send-event --send-metrics --send-service-check
```

## Docker Usage

This application is also available as a Docker image bundled with the latest version of Autorestic. Make sure to set `DD_DOGSTATSD_URL`

## Data Reported to Datadog

This tool collects and reports backup metrics from Autorestic to Datadog. Below is an overview of the data sent:

### Events

Events are sent to Datadog with each backup state (before, after, success, failure). The events contain the following data:

- **Title:** The event title will be set to the provided `--event-message`.
- **Text:** The event text includes information about the backup state, the Autorestic location, and the backend used (when applicable).
- **Tags:** Each event is tagged with `restic`, `autorestic`, `backup`, `app:<app_name>`, and `backend:<backend_name>` (for after, success, and failure states).
- **Alert Type:** The alert type is set according to the backup state: `info` for before and after, `error` for failure, and `success` for success.

### Metrics

Metrics are sent for the "after" and "success" backup states if `--send-metrics` is provided. The following metrics are collected:

- `snapshot_id`
- `parent_snapshot_id`
- `files_added`
- `files_changed`
- `files_unmodified`
- `dirs_added`
- `dirs_changed`
- `dirs_unmodified`
- `added_size`
- `processed_files`
- `processed_size`
- `processed_duration`
- `exit_code`

The metric name is in the format `restic.<metric_name>`. Each metric is a Gauge type and is tagged with `restic`, `autorestic`, `backup`, `app:<app_name>`, and `backend:<backend_name>`.

### Service Check

A service check is sent to Datadog if `--send-service-check` is provided. The service check data includes:

- **Name:** The service check name is set to `restic:<app_name>`.
- **Status:** The status is set to 0 (OK) for a success state, 1 (WARNING) for a before or after state, and 2 (CRITICAL) for a failure state.
- **Tags:** The service check is tagged with `restic`, `autorestic`, `backup`, `app:<app_name>`, and `backend:<backend_name>` (for after, success, and failure states).

Please note that the actual metrics sent to Datadog are dependent on the backup state and the environment variables available from Autorestic at the time the tool is run.

## Feedback and Contributions

Your feedback is welcome! If you have any questions, issues, or suggestions, feel free to open an issue in this repository. Contributions are also welcome. If you'd like to improve this application, please make a pull request. I appreciate your help!
