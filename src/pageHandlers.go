package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var isPrompt bool
var currentPromptIntent Intent

/*
*@replyIntent(): a function to run when /replyIntent is called
*@param w: ReponseWriter
*@param r: Request
 */
func replyIntent(w http.ResponseWriter, r *http.Request) {
	var messageReceived InputMesssage
	err := json.NewDecoder(r.Body).Decode(&messageReceived)
	//Check for error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// If the request is not prompt-related, query and answers normally
	if !isPrompt {
		for _, element := range listOfIntent {
			for _, e := range element.TrainingPhrases {
				if strings.Contains(strings.ToLower(messageReceived.MessageContent), strings.ToLower(e)) {
					if element.Prompt != nil {
						fmt.Fprintf(w, "%+v", element.Prompt[element.LAQ].PromptQuestion)
						isPrompt = true
						currentPromptIntent = element
						return
					}
					element.Reply.UserID = messageReceived.UserID
					fmt.Fprintf(w, "%+v", RMtoJson(element.Reply))
					return
				}
			}
		}
		//If the request has a prompt field, prompt the users for inputs.
	} else {
		if currentPromptIntent.LAQ < len(currentPromptIntent.Prompt)-1 {
			fmt.Fprintf(w, "%+v", PromptToJson(currentPromptIntent.Prompt[currentPromptIntent.LAQ+1]))
			currentPromptIntent.LAQ += 1
		} else {
			currentPromptIntent.Reply.UserID = messageReceived.UserID
			fmt.Fprintf(w, "%+v", RMtoJson(currentPromptIntent.Reply))
			resetPrompt()
		}
		return
	}
	/*If the bot doesnt understand what is going on, spit default message.*/
	defaultResponseMessage.UserID = messageReceived.UserID
	fmt.Fprintf(w, "%+v", RMtoJson(defaultResponseMessage))
}
