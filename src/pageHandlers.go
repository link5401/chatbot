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
	LAQ            int
}

type Intent struct {
	IntentName      string
	TrainingPhrases []string
	Reply           string
	Prompt          []Prompt
	LAQ             int
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
	LAQ:             0,
}
var defaultRegisterIntent = Intent{
	IntentName:      "Register",
	TrainingPhrases: []string{"Register", "I need Registration"},
	Reply:           "All good",
	Prompt:          []Prompt{NamePrompt, EmailPrompt},
	LAQ:             0,
}

var listOfIntent = []Intent{
	defaultHelloIntent,
	defaultRegisterIntent,
}

var isPrompt bool
var currentPromptIntent Intent

func replyIntent(w http.ResponseWriter, r *http.Request) {
	var messageReceived InputMesssage
	err := json.NewDecoder(r.Body).Decode(&messageReceived)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !isPrompt {
		for _, element := range listOfIntent {
			for _, e := range element.TrainingPhrases {
				if strings.Contains(strings.ToLower(messageReceived.MessageContent), strings.ToLower(e)) {
					if element.Prompt != nil {
						fmt.Fprintf(w, "%+v", element.Prompt[element.LAQ])
						isPrompt = true
						currentPromptIntent = element
						return
					}
					fmt.Fprintf(w, "%+v", element.Reply)
					return
				}
			}
		}
	} else if isPrompt {
		if currentPromptIntent.LAQ < len(currentPromptIntent.Prompt)-1 {
			fmt.Fprintf(w, "%+v", currentPromptIntent.Prompt[currentPromptIntent.LAQ+1].PromptQuestion)
			currentPromptIntent.LAQ += 1
		} else {
			fmt.Fprintf(w, "%+v", currentPromptIntent.Reply)
			isPrompt = false
			currentPromptIntent.LAQ = 0
			currentPromptIntent = Intent{}
		}
		return
	}
	fmt.Fprintf(w, "Sorry I don't understand what you are saying")

}
