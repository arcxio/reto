package elements

import (
	"github.com/arcxio/reto/internal/printer"
)

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
