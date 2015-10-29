package main

import (
	"bytes"
	"flag"
	"github.com/42wim/matterbridge/matterhook"
	"io"
	"os"
	"strings"
)

var (
	flagUserName, flagChannel, flagIconURL string
	flagPlainText, flagNoBuffer            bool
)

func init() {
	flag.StringVar(&flagUserName, "u", "mattertee", "This username is used for posting.")
	flag.StringVar(&flagChannel, "c", "", " Post input values to specified channel or user.")
	flag.StringVar(&flagIconURL, "i", "", "This url is used as icon for posting.")
	flag.BoolVar(&flagPlainText, "p", false, "Don't surround the post with triple backticks.")
	flag.BoolVar(&flagNoBuffer, "n", false, "Post input values without buffering.")
	flag.Parse()
}

func md(text string) string {
	return "```" + text + "```"
}

func main() {
	url := os.Getenv("MM_HOOK")
	m := matterhook.New(url, matterhook.Config{DisableServer: true})
	msg := matterhook.OMessage{}
	msg.UserName = flagUserName
	msg.Channel = flagChannel
	msg.IconURL = flagIconURL
	buf := new(bytes.Buffer)
	io.Copy(buf, os.Stdin)
	msg.Text = md(buf.String())
	if flagPlainText {
		msg.Text = buf.String()
	}
	if flagNoBuffer {
		texts := strings.Split(buf.String(), "\n")
		for _, text := range texts {
			msg.Text = md(text)
			if flagPlainText {
				msg.Text = text
			}
			m.Send(msg)
		}
	} else {
		m.Send(msg)
	}
}
