package components

import (
	"image/color"
	"strconv"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type StatusBar struct {
	Message string
}

func NewStatusBar(ln, col int) *StatusBar {
	return &StatusBar{
		Message: "Ln " + strconv.Itoa(ln) + ", Col " + strconv.Itoa(col) + " | UTF-8 | Go",
	}
}

func (s *StatusBar) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEnd,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			inset := layout.UniformInset(unit.Dp(4))
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				label := material.Label(theme, unit.Sp(10), s.Message)
				label.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				label.Alignment = text.End
				return label.Layout(gtx)
			})
		}),
	)
}
