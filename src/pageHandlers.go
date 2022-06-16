package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Prompt struct {
	ParamName      string
	ParamType      string
	PromptQuestion string
}

type Intent struct {
	IntentName      string
	TrainingPhrases []string
	Reply           string
	Prompt          []Prompt
}
type InputMesssage struct {
	UserID         string
	MessageContent string
}

var NamePrompt = Prompt{
	ParamName:      "name",
	ParamType:      "string",
	PromptQuestion: "What is your name",
}
var EmailPrompt = Prompt{
	ParamName:      "email",
	ParamType:      "email",
	PromptQuestion: "What is your email",
}
var defaultHelloIntent = Intent{
	IntentName:      "Hello",
	TrainingPhrases: []string{"Hello help", "Hi"},
	Reply:           "Hi, how can I help you",
}
var defaultRegisterIntent = Intent{
	IntentName:      "Register",
	TrainingPhrases: []string{"Register", "I need Registration"},
	Reply:           "All good",
	Prompt:          []Prompt{NamePrompt, EmailPrompt},
}

var listOfIntent = []Intent{
	defaultHelloIntent,
	defaultRegisterIntent,
}
var promptingIntent Intent
var promptingIndex int64

func replyIntent(w http.ResponseWriter, r *http.Request) {
	var messageReceived InputMesssage
	err := json.NewDecoder(r.Body).Decode(&messageReceived)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, element := range listOfIntent {
		for _, e := range element.TrainingPhrases {
			if strings.Contains(strings.ToLower(messageReceived.MessageContent), strings.ToLower(e)) {
				if element.Prompt != nil {
					promptingIntent = element
					fmt.Fprintf(w, "%+v", element.Prompt[promptingIndex])
					promptingIndex++
				} else {
					fmt.Fprintf(w, "%+v", element.Reply)
					return
				}
			}
		}
	}
	fmt.Fprintf(w, "Sorry I don't understand what you are saying")
}
