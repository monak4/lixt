package main

import (
	"image/color"
	"log"
	"os"
	"strings"

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
		window.Option(app.Title("Lixt"))

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
		LineHeight: unit.Sp(18),
	}
	backgroundColor := color.NRGBA{
		R: 30, G: 30, B: 30, A: 255,
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

			lines := strings.Split(editor.Text(), "\n")

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							// TODO: スクロール無効化 (Listを使うのをやめる？)
							//		 エディターの行と同期
							return material.List(theme, &list).Layout(gtx, len(lines), func(gtx layout.Context, index int) layout.Dimensions {
								list.Axis = layout.Vertical
								lbl := material.Label(theme, unit.Sp(12), LineNum(index+1))
								lbl.Color = white
								return layout.Inset{Left: unit.Dp(8), Top: unit.Dp(2)}.Layout(gtx, lbl.Layout)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							ed := material.Editor(theme, editor, "Type here...")
							ed.Color = white
							ed.HintColor = white_
							return layout.UniformInset(unit.Dp(8)).Layout(gtx, ed.Layout)
						}),
					)
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}

func LineNum(num int) string {
	if num < 10 {
		return " " + string('0'+rune(num))
	}
	return string('0'+rune(num/10)) + string('0'+rune(num%10))
}
