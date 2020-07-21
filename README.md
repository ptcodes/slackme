# Slackme

This is a very simple command-line tool for sending notifications to Slack.

It's useful when you want to be notified on a long-running task completion.

For example, we want to create a dump of a large database and be notified when it's done:

```
pg_dump -Fc customers_production > customers.sql && slackme database export is complete
```

We will get the following notification on completion:

[13:34:43 customer-db-server] database export is complete

Which also includes server time and hostname.

## Getting Started

* Create a new [Slack app](https://api.slack.com/apps) in the desired workspace
* Activate Incoming Webhooks and click on the Add New Webhook to Workspace button
* Select a Slack channel where notifications will be posted to
* Finally copy the Webhook URL and use it in the environment variable called SLACKME_WEBHOOK_URL

## Usage

```
slackme this is a test message
curl -O http://example.org/large_file.zip && slackme download is complete
./long_running_task.sh | xargs slackme
make build && slackme build is complete
```

