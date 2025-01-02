package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/arcxio/reto/internal/printer"
)

func doneFunc(textView *tview.TextView, p *printer.Printer, pages *Pages) func(key tcell.Key) {
	return func(key tcell.Key) {
		selection := textView.GetHighlights()
		switch key {
		case tcell.KeyTab:
			if len(selection) > 0 {
				i, _ := strconv.Atoi(selection[0])
				n := p.LinkCount()
				textView.Highlight(strconv.Itoa((i + 1) % n)).ScrollToHighlight()
				if arg := p.LinkUrl((i + 1) % n); arg != nil {
					pages.logView.SetText("selected link to " + *arg)
				}
			} else {
				textView.Highlight("0").ScrollToHighlight()
				if arg := p.LinkUrl(0); arg != nil {
					pages.logView.SetText("selected link to " + *arg)
				}
			}
		case tcell.KeyBacktab:
			n := p.LinkCount()
			if len(selection) > 0 {
				i, _ := strconv.Atoi(selection[0])
				textView.Highlight(strconv.Itoa((i - 1 + n) % n)).ScrollToHighlight()
				if arg := p.LinkUrl((i - 1 + n) % n); arg != nil {
					pages.logView.SetText("selected link to " + *arg)
				}
			} else {
				textView.Highlight(strconv.Itoa(n - 1)).ScrollToHighlight()
				if arg := p.LinkUrl(n - 1); arg != nil {
					pages.logView.SetText("selected link to " + *arg)
				}
			}
		case tcell.KeyEscape:
			if len(selection) > 0 {
				textView.Highlight()
				pages.logView.SetText("")
			}
		case tcell.KeyEnter:
			if len(selection) > 0 {
				i, _ := strconv.Atoi(selection[0])
				if arg := p.LinkUrl(i); arg != nil {
					go func(pages *Pages, arg string) {
						if err := pages.Open(arg); err != nil {
							pages.logView.SetText(err.Error())
						} else {
							pages.logView.SetText("")
						}
					}(pages, *arg)
				}
			}
		}
	}
}

func NewTextView(p *printer.Printer, pages *Pages) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		for cx := x; cx < x+width; cx++ {
			screen.SetContent(cx, y, tview.BoxDrawingsLightHorizontal, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
			screen.SetContent(cx, y+height-1, tview.BoxDrawingsLightHorizontal, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}
		return x, y + 1, width, height - 2
	})
	return textView.
		SetDynamicColors(true).
		SetRegions(true).
		SetDoneFunc(doneFunc(textView, p, pages))
}
