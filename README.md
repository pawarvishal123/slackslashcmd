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
      "name": "channel_id",
      "type": "string"
    },
    {
      "name": "channel_name",
      "type": "string"
    },
    {
      "name": "command",
      "type": "string"
    },
    {
      "name": "team_domain",
      "type": "string"
    },
    {
      "name": "team_id",
      "type": "string"
    },
    {
      "name": "text",
      "type": "string"
    },
    {
      "name": "user_id",
      "type": "string"
    },
    {
      "name": "user_name",
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
| Output       | Description                                                                                                                   |
|:-------------|:------------------------------------------------------------------------------------------------------------------------------|
| channel_id   | The unique ID of the channel the command came from                                                                            |
| channel_name | The name of the channel the command came from                                                                                 |
| command      | The Slack command that was executed                                                                                           |
| team_domain  | The name of the Slack workspace (if the URL of your Slack workspace is `myteam.slack.com`, the team_domain would be `myteam`) |
| team_id      | The unique ID of your Slack workspace                                                                                         |
| text         | The additional text that was sent                                                                                             |
| user_id      | The ID of the user that sent the message                                                                                      |
| user_name    | The name of the user that sent the message                                                                                    |

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
