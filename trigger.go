package slackslashcmd

import (
	"context"
	"fmt"
	"encoding/json"
	"flag"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/nlopes/slack"
)

var flogolog = logger.GetLogger("trigger-flogo-slackslashcmd")

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
	flogolog.Debugf("Initializing slack recv trigger...")
	t.handlers = ctx.GetHandlers()
	return nil
}

// Start implements ext.Trigger.Start
func (t *SlackSlashCmdTrigger) Start() error {

	fmt.Printf("Starting slack slash CMD trigger..")
	handlers := t.handlers
	
	flogolog.Debug("Processing handlers")
	for _, handler := range handlers {

		accessToken := handler.GetStringSetting("AccessToken")
		port := handler.GetStringSetting("Port")
		//accessToken := t.config.GetSetting("AccessToken")
		//flogolog.Debug("AccessToken: ", accessToken)
		var (
			verificationToken string
		)
	
		flag.StringVar(&verificationToken, "token", accessToken, accessToken)
		flag.Parse()
	
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			s, err := slack.SlashCommandParse(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
	
			if !s.ValidateToken(verificationToken) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			
			//params := &slack.Msg{Text: s.Text}
			t.RunHandler(handler, s.Command, s.Text, w)
		})
		fmt.Println("[INFO] Server listening")
		http.ListenAndServe(":"+port, nil)
		//log.Debugf("Processing Handler: %s", handler.ActionId)
	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *SlackSlashCmdTrigger) Stop() error {
	
	fmt.Printf("Stopping slack slash cmd server...")

	return nil
}

// RunHandler action on new Slack RTM message
func (t *SlackSlashCmdTrigger) RunHandler(handler *trigger.Handler, command string, params string, w http.ResponseWriter) {

	trgData := make(map[string]interface{})
	trgData["command"] = command
	trgData["params"] = params

	results, err := handler.Handle(context.Background(), trgData)

	if err != nil {
		fmt.Printf("Error starting action: ", err.Error())
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
			fmt.Printf(err.Error())
		}
		return
	}


	fmt.Printf("Ran Handler: [%s]", handler)
	
}
