package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	white  = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	white_ = color.NRGBA{R: 255, G: 255, B: 255, A: 80}
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	editor := &widget.Editor{
		SingleLine: false,
		Submit:     false,
		ReadOnly:   false,
		WrapPolicy: text.WrapWords,
	}
	backgroundColor := color.NRGBA{
		R: 30,
		G: 30,
		B: 30,
		A: 255,
	}

	var list widget.List
	list.Axis = layout.Horizontal
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			gtx.Constraints.Min = e.Size
			paint.Fill(gtx.Ops, backgroundColor)

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					ed := material.Editor(theme, editor, "Type here...")
					ed.Color = white
					ed.HintColor = white_

					return material.List(theme, &list).Layout(gtx, 1, func(gtx layout.Context, index int) layout.Dimensions {
						return layout.UniformInset(unit.Dp(8)).Layout(gtx, ed.Layout)
					})
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}
