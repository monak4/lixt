package ui

import (
	"strconv"

	"gioui.org/layout"
	"gioui.org/widget/material"

	"lixt/internals/ui/components"
)

type Window struct {
	theme     *material.Theme
	editor    *components.TextEditor
	statusBar *components.StatusBar
}

func NewWindow() *Window {
	return &Window{
		theme:     material.NewTheme(),
		editor:    components.NewTextEditor(),
		statusBar: components.NewStatusBar(1, 1),
	}
}

func (w *Window) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// TODO: エディターの行数表示
					// lines := strings.Split(editor.Text(), "\n")
					// TODO: スクロール無効化 (Listを使うのをやめる？)
					// エディターの行と同期
					// return material.List(theme, &list).Layout(gtx, len(lines), func(gtx layout.Context, index int) layout.Dimensions {
					// list.Axis = layout.Vertical
					// lbl := material.Label(theme, unit.Sp(12), LineNum(index+1))
					// lbl.Color = white
					// return layout.Inset{Left: unit.Dp(8), Top: unit.Dp(2)}.Layout(gtx, lbl.Layout)
					// })
					return layout.Dimensions{}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return w.editor.Layout(gtx, w.theme)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			ed := w.editor.ReturnEditor()
			ln, col := ed.CaretPos()
			w.statusBar.Message = "Ln " + strconv.Itoa(ln) + ", Col " + strconv.Itoa(col) + " | UTF-8 | Go"
			return w.statusBar.Layout(gtx, w.theme)
		}),
	)
}
