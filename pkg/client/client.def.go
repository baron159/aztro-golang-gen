package client

import (
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"repos.baron.solutions/seb/aztro/pkg/msg"
)

type AztroClientState struct {
	messages []*msg.Message
	input    *widget.Editor
	send     *widget.Clickable
	shaper   *text.Shaper
	font     text.Font
	fontSize unit.Sp
	text     string
}

// func A
func NewTestModel() *AztroClientState {
	return &AztroClientState{
		messages: []*msg.Message{
			msg.NewMessage("Hello", true),
			msg.NewMessage("Hi there!", false),
			msg.NewMessage("How are you doing today?", true),
			msg.NewMessage("I'm doing well, thanks for asking.", false),
			msg.NewMessage("That's great to hear!", true),
			msg.NewMessage("Yes, it is.", false),
		},
	}
}
