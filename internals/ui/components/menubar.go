package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MenuItem struct {
	Label    string
	Action   func()
	Shortcut string
}

type Menu struct {
	Label    string
	Items    []MenuItem
	Button   widget.Clickable
	IsOpen   bool
	ItemBtns []widget.Clickable
}

type MenuBar struct {
	Menus    []*Menu
	Theme    *material.Theme
	Active   *Menu
	Backdrop widget.Clickable
}

func NewMenuBar(theme *material.Theme) *MenuBar {
	return &MenuBar{
		Theme: theme,
	}
}

func (m *MenuBar) Layout(gtx layout.Context) layout.Dimensions {
	if m.Backdrop.Clicked(gtx) {
		m.closeAllMenus()
	}

	if m.Active != nil {
		m.drawBackdrop(gtx)
	}

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.layoutMenus(gtx)
		}),
	)
}

func (m *MenuBar) layoutMenus(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{}
}

func (m *MenuBar) layoutMenu(gtx layout.Context, menu *Menu) layout.Dimensions {
	return layout.Dimensions{}
}

func (m *MenuBar) layoutDropdown(gtx layout.Context, menu *Menu, btnDims layout.Dimensions) layout.Dimensions {
	return layout.Dimensions{}
}

func (m *MenuBar) drawBackdrop(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{}
}

func (m *MenuBar) closeAllMenus() {
	m.Active = nil
}
