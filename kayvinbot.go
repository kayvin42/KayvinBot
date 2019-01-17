package main

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "strings"

  "github.com/nlopes/slack"
)

func getenv(name string) string {
  v := os.Getenv(name)
  if v == "" {
    panic("missing required environment variable " + name)
  }
  return v
}

func main() {
  token := getenv("SLACKTOKEN")
  api := slack.New(token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "kayvin-bot: ", log.Lshortfile|log.LstdFlags)),
	)
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
      // Ignore hello


    case *slack.MessageEvent:
        info := rtm.GetInfo()
        text := ev.Text
        text = strings.TrimSpace(text)
        text = strings.ToLower(text)
        matched, _ := regexp.MatchString("test", text)
        fmt.Printf("Message: %v\n", text)

        if ev.User != info.User.ID && matched {
          rtm.SendMessage(rtm.NewOutgoingMessage("test response", ev.Channel))
        }

    case *slack.PresenceChangeEvent:
  			fmt.Printf("Presence Change: %v\n", ev)

  	case *slack.LatencyReport:
  			fmt.Printf("Current latency: %v\n", ev.Value)

    case *slack.RTMError:
        fmt.Printf("Error: %s\n", ev.Error())

    case *slack.InvalidAuthEvent:
        fmt.Printf("Invalid credentials")
        return

    default:
        // Take no action
      }
    }
}
