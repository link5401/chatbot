package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Person struct {
	Name string
	Age  int
}
type Intent struct {
	IntentName      string
	TrainingPhrases []string
	Reply           string
}
type InputMesssage struct {
	UserID         string
	MessageContent string
}

var defaultHelloIntent = Intent{
	IntentName:      "Hello",
	TrainingPhrases: []string{"Hello help", "Hi"},
	Reply:           "Hi, how can I help you",
}
var defaultHelpIntent = Intent{
	IntentName:      "Help",
	TrainingPhrases: []string{"Help", "I need help"},
	Reply:           "Haha, no",
}

// func createIntent(w http.ResponseWriter, r *http.Request) {
// 	err := json.NewDecoder(r.Body).Decode(&newIntent)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Fprintf(w, "Intent Created : %+v", newIntent)
// }
var listOfIntent = []Intent{
	defaultHelloIntent,
	defaultHelpIntent,
}

func replyIntent(w http.ResponseWriter, r *http.Request) {

	var messageReceived InputMesssage
	err := json.NewDecoder(r.Body).Decode(&messageReceived)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var flagF int = 0
	for _, element := range listOfIntent {
		for _, e := range element.TrainingPhrases {
			if strings.Contains(strings.ToLower(messageReceived.MessageContent), strings.ToLower(e)) {
				fmt.Fprintf(w, "%+v", element.Reply)
				flagF = 1
				break
			}
		}
		if flagF == 1 {
			break
		}
	}

}

// func personCreate(w http.ResponseWriter, r *http.Request) {
// 	// Declare a new Person struct.
// 	var p Person

// 	// Try to decode the request body into the struct. If there is an error,
// 	// respond to the client with the error message and a 400 status code.
// 	err := json.NewDecoder(r.Body).Decode(&p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Do something with the Person struct...
// 	fmt.Fprintf(w, "Person: %+v", p)
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello"))
// }
