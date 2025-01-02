package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/arcxio/reto/internal/elements"
	"github.com/arcxio/reto/internal/printer"
)

type page struct {
	title   string
	address string
}

type Pages struct {
	*tview.Pages

	pages       []page
	addressView *tview.TextView
	logView     *tview.TextView
}

func (p *Pages) Open(arg string) error {
	address, body, err := Get(arg)
	if err != nil {
		return err
	}
	defer body.Close()
	p.addressView.SetText(address)
	if name, item := p.GetFrontPage(); item != nil {
		front, _ := strconv.Atoi(name)
		p.pages = p.pages[:front+1]
		for i := front + 1; i < p.GetPageCount(); i++ {
			p.RemovePage(strconv.Itoa(i))
		}
	}
	pr := printer.NewPrinter(printer.TviewStyle)
	textView := NewTextView(pr, p)
	name := strconv.Itoa(p.GetPageCount())
	p.AddAndSwitchToPage(name, textView, true)
	if title, err := Print(body, textView, pr); err != nil {
		p.RemovePage(name)
		return err
	} else {
		p.pages = append(p.pages, page{title: title, address: address})
		return nil
	}
}

func Print(r io.Reader, w io.Writer, p *printer.Printer) (title string, err error) {
	p.SetTitle("")
	var isBlock, blockEnded bool
	elementStack := make([]elements.Element, 0)
	z := html.NewTokenizer(r)
	for {
		switch z.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()
			p.PushToken(token)
			element, isBlock := elements.FromAtom(token.DataAtom, p)
			if isBlock {
				p.IndentBlock(w)
				if p.InAtom(atom.Dd) && token.DataAtom != atom.Dt {
					p.IndentDefinition(w)
				}
			}
			if err := p.Print(w, element.Opening()); err != nil {
				return title, err
			}
			if token.Type == html.SelfClosingTagToken {
				p.PopToken()
				blockEnded = isBlock
			} else {
				elementStack = append(elementStack, element)
				blockEnded = false
			}
		case html.TextToken:
			if p.InAtom(atom.Style, atom.Script) {
				continue
			}
			text := string(z.Text())
			if !p.InAtom(atom.Pre) {
				cutset := "\n\r\t"
				text = strings.ReplaceAll(strings.Trim(text, cutset), "\n", " ")
			}
			if strings.ReplaceAll(text, " ", "") == "" {
				continue
			}
			if p.InAtom(atom.Title) {
				p.SetTitle(text)
				title = text
				continue
			}
			if blockEnded {
				p.IndentBlock(w)
			}
			if p.Style == printer.TviewStyle {
				text = tview.Escape(text)
			}
			if len(elementStack) > 0 {
				lastElement := elementStack[len(elementStack)-1]
				text = lastElement.Content(text)
			}
			if err := p.Print(w, text); err != nil {
				return title, err
			}
		case html.EndTagToken:
			lastElement := elementStack[len(elementStack)-1]
			if err := p.Print(w, lastElement.Closing()); err != nil {
				return title, err
			}
			p.PopToken()
			elementStack = elementStack[:len(elementStack)-1]
			blockEnded = isBlock
		case html.ErrorToken:
			err = p.PrintLinks(w)
			return
		}
	}
}

func (p *Pages) inputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'H':
			if name, _ := p.GetFrontPage(); name != "" && name != "0" {
				if i, err := strconv.Atoi(name); err == nil {
					p.SwitchToPage(strconv.Itoa(i - 1))
					printer.SetTitle(p.pages[i-1].title)
					p.addressView.SetText(p.pages[i-1].address)
				}
			}
			return nil
		case 'L':
			if name, _ := p.GetFrontPage(); name != "" {
				if i, err := strconv.Atoi(name); err == nil && i < p.GetPageCount()-1 {
					p.SwitchToPage(strconv.Itoa(i + 1))
					printer.SetTitle(p.pages[i+1].title)
					p.addressView.SetText(p.pages[i+1].address)
				}
			}
			return nil
		}
		return event
	}
}

func NewPages() *Pages {
	addressView := tview.NewTextView().SetWrap(false).SetTextAlign(tview.AlignCenter)
	logView := tview.NewTextView().SetWrap(false)

	p := &Pages{
		Pages:       tview.NewPages(),
		pages:       []page{},
		addressView: addressView,
		logView:     logView,
	}
	p.SetInputCapture(p.inputCapture())
	return p
}
