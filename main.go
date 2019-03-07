//////////////////////////////////////////////////
// Main Go File for the MathyBot Project.
// Created and Maintained by Aakash Sudhakar.
//////////////////////////////////////////////////

// Main package for general Golang functionality
package main

// Global imports, including Slack and Wolfram API
import (
	"os"

	"github.com/nlopes/slack"
)

// Main run function
func main() {
	slackClient := slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
}
