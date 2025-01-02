package elements

import (
	"github.com/arcxio/reto/internal/printer"
)

type Emphasis struct {
	p *printer.Printer
}

func (el Emphasis) Opening() string {
	switch el.p.Style {
	case printer.AnsiStyle:
		return "\033[3m_"
	case printer.TviewStyle:
		return "[::u]_"
	default:
		return "_"
	}
}

func (el Emphasis) Content(text string) string {
	switch el.p.Style {
	case printer.AnsiStyle:
		el.p.PushFormatter("\033[3m")
	case printer.TviewStyle:
		el.p.PushFormatter("[::u]")
	}
	return text
}

func (el Emphasis) Closing() string {
	switch el.p.Style {
	case printer.AnsiStyle:
		el.p.PopFormatter()
		return "_\033[24m"
	case printer.TviewStyle:
		el.p.PopFormatter()
		return "_[::U]"
	default:
		return "_"
	}
}

type Strong struct {
	p *printer.Printer
}

func (el Strong) Opening() string {
	switch el.p.Style {
	case printer.AnsiStyle:
		return "\033[1m*"
	case printer.TviewStyle:
		return "[::b]*"
	default:
		return "*"
	}
}

func (el Strong) Content(text string) string {
	switch el.p.Style {
	case printer.AnsiStyle:
		el.p.PushFormatter("\033[1m")
	case printer.TviewStyle:
		el.p.PushFormatter("[::b]")
	}
	return text
}

func (el Strong) Closing() string {
	switch el.p.Style {
	case printer.AnsiStyle:
		el.p.PopFormatter()
		return "*\033[22m"
	case printer.TviewStyle:
		el.p.PopFormatter()
		return "*[::B]"
	default:
		return "*"
	}
}
