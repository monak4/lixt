package components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"lixt/internals/utility"
)

type TextEditor struct {
	editor         widget.Editor
	bracketHandler *utility.BracketHandler
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
		bracketHandler: utility.NewBracketsHandler(),
	}
}

var (
	white  = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	white_ = color.NRGBA{R: 255, G: 255, B: 255, A: 80}
)

func (t *TextEditor) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	ed := material.Editor(theme, &t.editor, "Type here...")
	ed.Color = white
	ed.HintColor = white_

	var list widget.List
	list.Axis = layout.Horizontal
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.List(theme, &list).Layout(gtx, 1, func(gtx layout.Context, index int) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, ed.Layout)
			})
		}),
	)
}

func (t *TextEditor) ReturnEditor() widget.Editor {
	return t.editor
}
