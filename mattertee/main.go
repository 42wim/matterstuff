package main

import (
	"bytes"
	"github.com/42wim/matterbridge/matterhook"
	"io"
	"os"
)

func main() {
	var title string
	url := os.Getenv("MM_HOOK")
	m := matterhook.New(url, matterhook.Config{})
	msg := matterhook.OMessage{}
	msg.UserName = "mattertee"
	buf := new(bytes.Buffer)
	io.Copy(buf, os.Stdin)
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	msg.Text = title + "\n"
	msg.Text = msg.Text + "```" + buf.String() + "```"
	m.Send(msg)
}
