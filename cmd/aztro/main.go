package main

import (
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func (m *model) update(gtx layout.Context) {
	// update the message input field
	m.input.SetText(m.text)

	// handle message sending
	for m.send.Clicked() {
		if m.text != "" {
			m.messages = append(m.messages, message{Content: m.text, Sent: true})
			m.text = ""
		}
	}

	// draw the messages
	drawMessages(gtx, th, m.messages, "User")

	layout.Flex{Alignment: layout.Start}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.input.Layout(gtx, m.shaper, m.font, m.fontSize, func(gtx layout.Context, lt *text.Layout, font text.Font, size unit.Value, content layout.Widget) layout.Dimensions {
					return layout.Flex{Alignment: layout.End}.Layout(gtx,
						layout.Flexed(1, content),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, m.send, "Send").Layout(gtx)
						}),
					)
				})
			})
		}),
	)
}

// func C
func drawMessages(gtx layout.Context, th *material.Theme, msgs []message, sender string) layout.Dimensions {
	return layout.Flex{Direction: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{}
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
								return layout.Dimensions{}
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

// function D
func messageBubble(th *material.Theme, body string, isSent bool) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		var (
			background color.NRGBA
			triangle   Path
			maxWidth   int
			padding    = unit.Dp(8)
			inset      = layout.Inset{Top: padding, Bottom: padding}
			direction  = layout.W
		)

		if isSent {
			background = sentColor
			triangle = triangleRight
			direction = layout.E
		} else {
			background = receivedColor
			triangle = triangleLeft
		}

		dims := material.Body1(th, body).Layout(gtx)

		width := dims.Size.X + padding.Dp*2
		if width > maxWidth {
			maxWidth = width
		}

		return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, background, triangle.Op(gtx.Ops))

			return layout.Stack{Alignment: layout.W}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					return dims
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					square := f32.Rectangle{
						Max: layout.FPt(gtx.Constraints.Max),
					}
					clip.RRect{
						Rect: square,
						NE:   6, NW: 6, SE: 6, SW: 6,
					}.Add(gtx.Ops)
					paint.Fill(gtx.Ops, text.Color.Black)
					return layout.Dimensions{}
				}),
			)
		})
	}
}

// func C
// func drawMessages(gtx layout.Context, th *material.Theme, msgs []message, sender string) layout.Dimensions {
// 	return layout.Flex{Direction: layout.Vertical}.Layout(gtx,
// 		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 			return layout.Dimensions{}
// 		}),
// 		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 				return layout.Flex{Direction: layout.Vertical}.Layout(gtx,
// 					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 						return layout.Dimensions{}
// 					}),
// 					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 						return layout.Flex{}.Layout(gtx,
// 							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 								return layout.Dimensions{}
// 							}),
// 							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
// 								return layout.Dimensions{}
// 							}),
// 							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 								return layout.Dimensions{}
// 							}),
// 						)
// 					}),
// 				)
// 			})
// 		}),
// 	)
// }
