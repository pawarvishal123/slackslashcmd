# Slack Slash Command Trigger

This trigger provides your flogo application the ability to react to Slack "Slash Commands".

## Installation

```bash
flogo install github.com/pawarvishal123/slackslashcmd
```
Link for flogo web:
```
https://github.com/github.com/pawarvishal123/slackslashcmd
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
      "name": "accessToken",
      "type": "string",
	  "required":"true"
    },
    {
      "name": "port",
      "type": "string",
	  "required":"true"
    }]
```

## Outputs
| Output  | Description                              |
|:--------|:-----------------------------------------|
| command | The Slack command that was executed      |
| params  | The additional parameters that were sent |

## Reply
| Reply | Description                        |
|:------|:-----------------------------------|
| data  | The response you want to send back |

## Handler
| Handler     | Description                                                                                                       |
|:------------|:------------------------------------------------------------------------------------------------------------------|
| accessToken | The Verification Token of your Slack app. The trigger uses this to validate the message actually comes from Slack |
| port        | The HTTP port your app will listen on                                                                             |

## Example Configurations
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
