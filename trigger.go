package slackslashcmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/nlopes/slack"
)

var log = logger.GetLogger("trigger-flogo-slackslashcmd")

// SlackSlashCmdTrigger is Slack slash cmd trigger
type SlackSlashCmdTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &SlackSlashCmdFactory{metadata: md}
}

// SlackSlashCmdFactory Trigger factory
type SlackSlashCmdFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *SlackSlashCmdFactory) New(config *trigger.Config) trigger.Trigger {
	return &SlackSlashCmdTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *SlackSlashCmdTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize implements trigger.Init
func (t *SlackSlashCmdTrigger) Initialize(ctx trigger.InitContext) error {
	log.Debugf("Initializing Slack Slash Command trigger...")
	t.handlers = ctx.GetHandlers()
	return nil
}

// Start implements ext.Trigger.Start
func (t *SlackSlashCmdTrigger) Start() error {
	log.Infof("Starting Slack Slash Command trigger...")
	handlers := t.handlers

	log.Debugf("Processing handlers")
	for _, handler := range handlers {

		accessToken := handler.GetStringSetting("accessToken")
		port := handler.GetStringSetting("port")

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			s, err := slack.SlashCommandParse(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if !s.ValidateToken(accessToken) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			t.RunHandler(handler, s.ChannelID, s.ChannelName, s.Command, s.TeamDomain, s.TeamID, s.Text, s.UserID, s.UserName, w)
		})
		log.Infof("Server listening on port [%s]", port)
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *SlackSlashCmdTrigger) Stop() error {
	log.Infof("Stopping Slack Slash Command triggers...")
	return nil
}

// RunHandler action on new Slack message
func (t *SlackSlashCmdTrigger) RunHandler(handler *trigger.Handler, channelID string, channelName string, command string, teamDomain string, teamID string, text string, userID string, userName string, w http.ResponseWriter) {
	trgData := make(map[string]interface{})
	trgData["channel_id"] = channelID
	trgData["channel_name"] = channelName
	trgData["command"] = command
	trgData["team_domain"] = teamDomain
	trgData["team_id"] = teamID
	trgData["text"] = text
	trgData["user_id"] = userID
	trgData["user_name"] = userName

	results, err := handler.Handle(context.Background(), trgData)

	if err != nil {
		log.Errorf("Error starting action: %s", err.Error())
	}

	var replyData interface{}

	if len(results) != 0 {
		dataAttr, ok := results["data"]
		if ok {
			replyData = dataAttr.Value()
		}
	}

	if replyData != nil {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(replyData); err != nil {
			log.Errorf("Error replying to Slack: %s", err.Error())
		}
		return
	}

	log.Debugf("Ran Handler: [%s]", handler)
}
