package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"

	"lixt/internals/ui"
)

var (
	backgroundColor = color.NRGBA{
		R: 30, G: 30, B: 30, A: 255,
	}
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
	w := ui.NewWindow()

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			gtx.Constraints.Min = e.Size
			paint.Fill(gtx.Ops, backgroundColor)

			w.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

// TODO: 行数表示用のファイルを作る
// LineNum は行数表示をフォーマットするための関数
func LineNum(num int) string {
	if num < 10 {
		return " " + string('0'+rune(num))
	}
	return string('0'+rune(num/10)) + string('0'+rune(num%10))
}
