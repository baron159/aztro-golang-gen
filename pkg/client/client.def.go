package client

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"repos.baron.solutions/seb/aztro/pkg/msg"
)

type AztroClientState struct {
	Messages      []*msg.Message
	Input         *widget.Editor
	Send          *widget.Clickable
	Shaper        *text.Shaper
	Font          text.Font
	FontSize      unit.Sp
	SelectedParty interface{}
	Text          string
}

// func A
func NewTestModel() *AztroClientState {
	return &AztroClientState{
		Messages: []*msg.Message{
			msg.NewMessage("Hello", true),
			msg.NewMessage("Hi there!", false),
			msg.NewMessage("How are you doing today?", true),
			msg.NewMessage("I'm doing well, thanks for asking.", false),
			msg.NewMessage("That's great to hear!", true),
			msg.NewMessage("Yes, it is.", false),
		},
	}
}

// update updates the state of the Aztro client.
func (a *AztroClientState) Update(gtx layout.Context) {
	// Update message input field
	a.Input.SetText(a.Text)

	// Handle message sending
	for a.Send.Clicked() {
		if a.Text != "" {
			// Append new message to the messages slice
			a.Messages = append(a.Messages, &msg.Message{Party: "user", Content: a.Text, Sent: true})
			a.Text = ""
		}
	}

	// Draw messages
	messages := layout.Flex{Alignment: layout.Start}
	layout.Flex{Alignment: layout.Start}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return a.Input.Layout(gtx, a.Shaper, a.Font, a.FontSize, func(gtx layout.Context, lt *text.Layout, font text.Font, size unit.Value, content layout.Widget) layout.Dimensions {
					return layout.Flex{Alignment: layout.End}.Layout(gtx,
						layout.Flexed(1, content),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return a.Send.Layout(gtx, material.Body1(a.Theme), "Send")
						}),
					)
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Alignment: layout.Start}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return messages.Layout(gtx, len(a.Messages), func(gtx layout.Context, i int) layout.Dimensions {
						return a.drawMessage(gtx, a.Messages[i])
					})
				}),
			)
		}),
	)
}

func drawMessages(gtx layout.Context, th *material.Theme, msgs []message, sender string) layout.Dimensions {
	inset := layout.Inset{
		Top:    unit.Dp(8),
		Bottom: unit.Dp(8),
	}
	return layout.Flex{Direction: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{}
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Direction: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Dimensions{}
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Direction: layout.Vertical}.Layout(gtx, len(msgs), func(gtx layout.Context, i int) layout.Dimensions {
									return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return messageBubble(gtx, msgs[i], sender, th)
									})
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Dimensions{}
							}),
						)
					}),
				)
			})
		}),
	)
}

func messageBubble(gtx layout.Context, msg msg.Message, sender string, th *material.Theme) layout.Dimensions {
	bubbleColor := th.Color.Text
	if msg.Sent {
		bubbleColor = th.Color.Primary
	}

	return material.Card(th, unit.Dp(4), func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Bottom: unit.Dp(4),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Direction: layout.Vertical,
						Alignment: layout.Start,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{}
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.H6(th, msg.Sender).Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							p := material.Body1(th, msg.Content)
							p.Color = bubbleColor
							return p.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{}
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{}
				}),
			)
		})
	})
}
