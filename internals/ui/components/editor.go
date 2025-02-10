package components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TextEditor struct {
	editor widget.Editor
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		editor: widget.Editor{
			SingleLine: false,
			Submit:     false,
			ReadOnly:   false,
			WrapPolicy: text.WrapWords,
			LineHeight: unit.Sp(18),
		},
	}
}

var (
	white  = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	white_ = color.NRGBA{R: 255, G: 255, B: 255, A: 80}
)

func (t *TextEditor) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			ed := material.Editor(theme, &t.editor, "Type here...")
			ed.Color = white
			ed.HintColor = white_

			return layout.UniformInset(unit.Dp(8)).Layout(gtx, ed.Layout)
		}),
	)
}
