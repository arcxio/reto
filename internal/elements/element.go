package elements

import (
	"golang.org/x/net/html/atom"

	"github.com/arcxio/reto/internal/printer"
)

type Element interface {
	Opening() string
	Content(text string) string
	Closing() string
}

var atomToElementCreator = map[atom.Atom]func(p *printer.Printer) Element{
	atom.A:      func(p *printer.Printer) Element { return &Link{p: p} },
	atom.B:      func(p *printer.Printer) Element { return &Strong{p} },
	atom.Strong: func(p *printer.Printer) Element { return &Strong{p} },
	atom.Br:     func(p *printer.Printer) Element { return &Wrapper{"\\\n", ""} },
	atom.Del:    func(p *printer.Printer) Element { return &Wrapper{"{-", "-}"} },
	atom.Dt:     func(p *printer.Printer) Element { return &Wrapper{": ", ""} },
	atom.Em:     func(p *printer.Printer) Element { return &Emphasis{p} },
	atom.I:      func(p *printer.Printer) Element { return &Emphasis{p} },
	atom.U:      func(p *printer.Printer) Element { return &Emphasis{p} },
	atom.H1:     func(p *printer.Printer) Element { return &Heading{p.Style, 1} },
	atom.H2:     func(p *printer.Printer) Element { return &Heading{p.Style, 2} },
	atom.H3:     func(p *printer.Printer) Element { return &Heading{p.Style, 3} },
	atom.H4:     func(p *printer.Printer) Element { return &Heading{p.Style, 4} },
	atom.H5:     func(p *printer.Printer) Element { return &Heading{p.Style, 5} },
	atom.H6:     func(p *printer.Printer) Element { return &Heading{p.Style, 6} },
	atom.Hr:     func(p *printer.Printer) Element { return &Wrapper{"****", ""} },
	atom.Ins:    func(p *printer.Printer) Element { return &Wrapper{"{+", "+}"} },
	atom.Li:     func(p *printer.Printer) Element { return &ListItem{p} },
	atom.Mark:   func(p *printer.Printer) Element { return &Wrapper{"{=", "=}"} },
	atom.Sub:    func(p *printer.Printer) Element { return &Wrapper{"~", "~"} },
	atom.Sup:    func(p *printer.Printer) Element { return &Wrapper{"^", "^"} },
}

var blockAtoms = map[atom.Atom]struct{}{
	atom.Address:    {},
	atom.Article:    {},
	atom.Aside:      {},
	atom.Blockquote: {},
	atom.Canvas:     {},
	atom.Dd:         {},
	atom.Div:        {},
	atom.Dl:         {},
	atom.Dt:         {},
	atom.Fieldset:   {},
	atom.Figcaption: {},
	atom.Figure:     {},
	atom.Footer:     {},
	atom.Form:       {},
	atom.H1:         {},
	atom.H2:         {},
	atom.H3:         {},
	atom.H4:         {},
	atom.H5:         {},
	atom.H6:         {},
	atom.Header:     {},
	atom.Hr:         {},
	atom.Li:         {},
	atom.Main:       {},
	atom.Nav:        {},
	atom.Noscript:   {},
	atom.Ol:         {},
	atom.P:          {},
	atom.Pre:        {},
	atom.Section:    {},
	atom.Table:      {},
	atom.Tfoot:      {},
	atom.Ul:         {},
	atom.Video:      {},
}

func FromAtom(a atom.Atom, p *printer.Printer) (el Element, isBlock bool) {
	_, isBlock = blockAtoms[a]
	if newElement, ok := atomToElementCreator[a]; ok {
		return newElement(p), isBlock
	} else {
		return Nop{}, isBlock
	}
}
