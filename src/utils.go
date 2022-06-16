package main

import (
	"encoding/json"
	"fmt"
)

type InputMesssage struct {
	UserID         string
	MessageContent string
}
type ResponseMessage struct {
	UserID         string
	MessageContent string
}
type Prompt struct {
	ParamName      string
	ParamType      string
	PromptQuestion string
	LAQ            int
}

type Intent struct {
	IntentName      string
	TrainingPhrases []string
	Reply           ResponseMessage
	Prompt          []Prompt
	LAQ             int
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
	TrainingPhrases: []string{"Hello", "Hi"},
	Reply:           ResponseMessage{MessageContent: "Hi how can I help you"},
	LAQ:             0,
}
var defaultRegisterIntent = Intent{
	IntentName:      "Register",
	TrainingPhrases: []string{"Register", "I need Registration"},
	Reply:           ResponseMessage{MessageContent: "All good"},
	Prompt:          []Prompt{NamePrompt, EmailPrompt},
	LAQ:             0,
}
var defaultResponseMessage = ResponseMessage{
	MessageContent: "Sorry I dont get what you are saying",
}
var listOfIntent = []Intent{
	defaultHelloIntent,
	defaultRegisterIntent,
}

/*
 * @resetPrompt(): reset parameters to their normal state when prompts run out, letting the bot return to normal mode
 */
func resetPrompt() {
	isPrompt = false
	currentPromptIntent.LAQ = 0
	currentPromptIntent = Intent{}
}

/*
 *Converting to JSON format utils */
/*
* @param  rm: A response message struct
* @param  p: A prompt struct
* @return a JSON format string.
 */
func RMtoJson(rm ResponseMessage) string {
	byteArray, err := json.Marshal(rm)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}
func PromptToJson(p Prompt) string {
	byteArray, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}
