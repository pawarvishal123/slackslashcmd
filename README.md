# Slack Slash Command Trigger
Flogo trigger activity for Slack Slash Commands. You need to register slack slash command for your "slack app" and provide public URL of you flogo app endpoint. 
Slack "Verification Token" should be provided as input for activity configuration.

## Installation

```bash
flogo install github.com/pawarvishal123/slackslashcmd
```

## Schema
Settings, Outputs and Endpoint:

```json
{
  "output": [
    {
      "name": "command",
      "type": "string"
    },
    {
      "name": "params",
      "type": "string"
    }
  ],
  "reply": [
    {
      "name": "data",
      "type": "any"
    }
  ],
  "handler": {
    "settings": [{
      "name": "AccessToken",
      "type": "string",
	  "required":"true"
    },
    {
      "name": "Port",
      "type": "string",
	  "required":"true"
    }]
```

## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the Slack Slash command Trigger.

### Start a flow
Provide access token and port number to receive command from slack channel. The access token should have scope and permissions configured.

```json
{
  "triggers": [
    {
      "id": "receive_slack_rtm_messages",
      "ref": "https://github.com/pawarvishal123/slackslashcmd",
      "name": "Slack Slash Command Trigger",
      "description": "Slack Slash Command Trigger",
      "settings": {},
      "handlers": [
        {
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "data": {
              "flowURI": "res://flow:test_trigger"
            }
          },
          "settings": {
            "AccessToken": "<<YOUR-SLACK-VERFIFICATION-TOKEN>>",
            "Port": "<<YOUR-PORT>>"
          }
      }
    ]
}
```

## Third Party Library
Slack API in Go - [https://github.com/nlopes/slack](https://github.com/nlopes/slack)
