package components

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// 状態管理
type MenuBar struct {
	fileBtn  widget.Clickable
	editBtn  widget.Clickable
	helpBtn  widget.Clickable
	fileMenu *widget.Bool
	editMenu *widget.Bool
	helpMenu *widget.Bool
}

func NewMenuBar() *MenuBar {
	return &MenuBar{
		fileMenu: new(widget.Bool),
		editMenu: new(widget.Bool),
		helpMenu: new(widget.Bool),
	}
}

func (m *MenuBar) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.layoutMenu(gtx, theme, &m.fileBtn, m.fileMenu, "File", []string{
				"New", "Open", "Save", "Save As", "Exit",
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.layoutMenu(gtx, theme, &m.editBtn, m.editMenu, "Edit", []string{
				"Undo", "Redo", "Cut", "Copy", "Paste",
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.layoutMenu(gtx, theme, &m.helpBtn, m.helpMenu, "Help", []string{
				"About", "Version",
			})
		}),
	)
}

func (m *MenuBar) layoutMenu(gtx layout.Context, theme *material.Theme,
	btn *widget.Clickable, open *widget.Bool, label string, items []string) layout.Dimensions {

	// if btn.Clicked(layout.Context{}) {
	// 	*open = !*open
	// }

	menu := layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return material.Button(theme, btn, label).Layout(gtx)
		}),
	)

	if *&open.Value {
		op.Offset(menu.Size).Add(gtx.Ops)
		layout.Stack{}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				list := layout.List{Axis: layout.Vertical}
				return list.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
					btn := widget.Clickable{}
					return material.Button(theme, &btn, items[i]).Layout(gtx)
				})
			}),
		)
	}

	// return menu
	return layout.Dimensions{}
}

// func (m *MenuBar) IsFileMenuOpen() bool { return *m.fileMenu }
// func (m *MenuBar) IsEditMenuOpen() bool { return *m.editMenu }
// func (m *MenuBar) IsHelpMenuOpen() bool { return *m.helpMenu }
