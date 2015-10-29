package main

import (
	"bytes"
	"flag"
	"github.com/42wim/matterbridge/matterhook"
	"io"
	"os"
)

var (
	flagUserName, flagChannel, flagIconURL string
	flagPlainText                          bool
)

func init() {
	flag.StringVar(&flagUserName, "u", "mattertee", "This username is used for posting.")
	flag.StringVar(&flagChannel, "c", "", " Post input values to specified channel or user.")
	flag.StringVar(&flagIconURL, "i", "", "This url is used as icon for posting.")
	flag.BoolVar(&flagPlainText, "p", false, "Don't surround the post with triple backticks.")
	flag.Parse()
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
	if flagPlainText {
		msg.Text = buf.String()
	} else {
		msg.Text = "```" + buf.String() + "```"
	}
	m.Send(msg)
}
