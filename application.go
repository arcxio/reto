package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewApplication(pages *Pages) *tview.Application {
	app := tview.NewApplication()
	draw := func() { app.Draw() }
	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		AddItem(pages.addressView.SetChangedFunc(draw), 0, 0, 1, 1, 0, 0, false).
		AddItem(pages, 1, 0, 1, 1, 0, 0, true).
		AddItem(pages.logView.SetChangedFunc(draw), 2, 0, 1, 1, 0, 0, false)
	return app.
		EnableMouse(true).
		SetRoot(grid, true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' {
				app.Stop()
				return nil
			}
			return event
		})
}
