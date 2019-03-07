//////////////////////////////////////////////////
// Main Go File for the MathyBot Project.
// Created and Maintained by Aakash Sudhakar.
//////////////////////////////////////////////////

// Main package for general Golang functionality
package main

// Global imports, including Slack and Wolfram API
import (
	"fmt" // Permits line printing to console/logs
	"os"  // Permits OS operations/functionality

	"github.com/nlopes/slack" // External Slack API
)

// Initializing the Slack Client API
var (
	slackClient *slack.Client
)

// Main run function
func main() {
	// Instantiating our Slack client to communicate with Make School's Slack
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))

	// Instantiating real-time messaging with our Slackbot
	realTimeMSG := slackClient.NewRTM()

	// Wrapping our RTM connection in a concurrent Go Routine
	go realTimeMSG.ManageConnection()

	// Check for real-time messages hitting the Slackbot
	for msg := range realTimeMSG.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			// TODO: Create separate Go Routine to handle Slack messages
		}
	}
}

// Global function for handling real-time messaging events via the Slackbot
func handleMSGEvent(event *slack.MessageEvent) {
	fmt.Printf("%v\n", event)
}
