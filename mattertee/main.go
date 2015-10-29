package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/42wim/matterbridge/matterhook"
	"io"
	"os"
	"os/user"
	"time"
)

var (
	flagUserName, flagChannel, flagIconURL, flagMatterURL string
	flagPlainText, flagNoBuffer, flagExtra                bool
)

func init() {
	flag.StringVar(&flagUserName, "u", "mattertee", "This username is used for posting.")
	flag.StringVar(&flagChannel, "c", "", " Post input values to specified channel or user.")
	flag.StringVar(&flagIconURL, "i", "", "This url is used as icon for posting.")
	flag.StringVar(&flagMatterURL, "m", "", "Mattermost incoming webhooks URL.")
	flag.BoolVar(&flagPlainText, "p", false, "Don't surround the post with triple backticks.")
	flag.BoolVar(&flagNoBuffer, "n", false, "Post input values without buffering.")
	flag.BoolVar(&flagExtra, "x", false, "Add extra info (user/hostname/timestamp).")
	flag.Parse()
}

func md(text string) string {
	return "```\n" + text + "```"
}

func extraInfo() string {
	u, _ := user.Current()
	hname, _ := os.Hostname()
	return "\n" + u.Username + "@" + hname + " (_" + time.Now().Format(time.RFC3339) + "_)\n"
}

func main() {
	url := os.Getenv("MM_HOOK")
	if flagMatterURL != "" {
		url = flagMatterURL
	}
	m := matterhook.New(url, matterhook.Config{DisableServer: true})
	msg := matterhook.OMessage{}
	msg.UserName = flagUserName
	msg.Channel = flagChannel
	msg.IconURL = flagIconURL
	if flagNoBuffer {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			msg.Text = md(line)
			if flagPlainText {
				msg.Text = line
			}
			m.Send(msg)
		}
	} else {
		buf := new(bytes.Buffer)
		io.Copy(buf, os.Stdin)
		text := buf.String()
		fmt.Print(text)
		msg.Text = md(text)
		if flagPlainText {
			msg.Text = text
		}
		if flagExtra {
			msg.Text += extraInfo()
		}
		m.Send(msg)
	}
}
