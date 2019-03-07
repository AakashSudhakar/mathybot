//////////////////////////////////////////////////
// Main Go File for the MathyBot Project.
// Created and Maintained by Aakash Sudhakar.
//////////////////////////////////////////////////

// Main package for general Golang functionality
package main

// Global imports, including Slack and Wolfram API
import (
	"log" // Permits console logging
	"os"  // Permits OS operations/functionality

	wit "github.com/christianrondeau/go-wit" // External Wit.ai API
	"github.com/nlopes/slack"                // External Slack API
	// External Wolfram API
)

// Global constant holding message entity ideal confidence threshold
const optimalEntityConfidenceThreshold = 0.5

// Initializing the client APIs
var (
	slackClient   *slack.Client
	witClient     *wit.Client
	wolframClient *wolfram.Client
)

// Main run function
func main() {
	// Setting our client APIs to communicate across Make School's Slack
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	witClient = wit.NewClient(os.Getenv("WIT_AI_ACCESS_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	// Instantiating real-time messaging with our Slackbot
	realTimeMSG := slackClient.NewRTM()

	// Wrapping our RTM connection in a concurrent Go Routine
	go realTimeMSG.ManageConnection()

	// Checking for real-time messages hitting the Slackbot
	for msg := range realTimeMSG.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(event.BotID) == 0 {
				// Handling real-time messaging event via Go Routine
				go handleMSGEvent(event)
			}
		}
	}
}

// Global function for handling real-time messaging events via the Slackbot
func handleMSGEvent(event *slack.MessageEvent) {
	// fmt.Printf("%v\n", event) 	// Quick & dirty debugger code
	textRTM := event.Msg.Text
	res, err := witClient.Message(textRTM)

	// Error handling for response retrieval failure
	if err != nil {
		log.Printf("MESSAGE HANDLING ERROR: Unable to get response from Wit.ai server.\nError Details: %v", err)
		return
	}

	// Initializing variables to hold ideal message characteristic entity for NLP
	var (
		optimalEntityKey string
		optimalEntity    wit.MessageEntity
	)

	// Mapping over all message entities to grab ideal entity for NLP based on highest confidence
	for entityKey, entityValueMap := range res.Entities {
		for _, entity := range entityValueMap {
			if (entity.Confidence > optimalEntityConfidenceThreshold) && (entity.Confidence > optimalEntity.Confidence) {
				optimalEntityKey = entityKey
				optimalEntity = entity
			}
		}
	}

	// Responding to user based on characterized ideal MSG entity
	sendUserResponse(event, optimalEntityKey, optimalEntity)
}

// Global function for sending replies to user based on RTM NLP characterization
func sendUserResponse(event *slack.MessageEvent, optimalEntityKey string, optimalEntity wit.MessageEntity) {
	switch optimalEntityKey {
	case "greetings":
		slackClient.PostMessage(
			event.User,
			slack.MsgOptionText("Hello! I am WolfBot and I am here to answer your questions.", false),
			slack.MsgOptionAsUser(true),
		)
		return
	case "wolfram_search_query":
		res, err := wolframClient.GetSpokentAnswerQuery(optimalEntity.Value.(string), wolfram.metric, 1000)
		if err != nil {
			slackClient.PostMessage()
		}
	}

	slackClient.PostMessage(
		event.User,
		slack.MsgOptionText("WARNING: User input is unclear. Try clarifying your question?", false),
		slack.MsgOptionAsUser(true),
	)
}
