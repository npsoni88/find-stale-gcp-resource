// logging to STDOUT instead of print functions
// containerize the code
// running this script (pod) once every day instead of all the time
//how does a container get GOOGLE_APPLICATION_CREDENTIALS
// service account across all of our projects
// If []projects should be a variable outside of this logic

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/npsoni88/gcp-go/disks/staleDisks"
	"github.com/slack-go/slack"
)

var (
	api = slack.New(os.Getenv("SLACK_TOKEN"))
)

func main() {
	ctx := context.Background()
	projects := []string{
		"project1",
		"project2",
	}

	for _, v := range projects {
		unusedDisks := staleDisks.GetUnusedDisks(ctx, v)
		numberOfDisks := strconv.Itoa(len((unusedDisks)))
		attachment := slack.Attachment{
			Pretext: "There are " + numberOfDisks + " unused Disks in " + v + " project",
			Text:    strings.ReplaceAll(strings.Join(unusedDisks, " "), " ", "\n"),
		}
		_, _, err := api.PostMessage("C02JWS1EPK4",
			slack.MsgOptionAttachments(attachment))
		if err != nil {
			fmt.Println("error while sending the message", err)
		}
		log.Printf("Published unused disks of %s project", v)
	}

}
