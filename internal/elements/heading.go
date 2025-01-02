package elements

import (
	"strings"

	"github.com/arcxio/reto/internal/printer"
)

type Heading struct {
	style printer.Style
	level int
}

func (el Heading) Opening() string {
	var sb strings.Builder
	switch el.style {
	case printer.AnsiStyle:
		sb.WriteString("\033[1m")
	case printer.TviewStyle:
		sb.WriteString("[::b]")
	}
	for i := 0; i < el.level; i++ {
		sb.WriteRune('#')
	}
	sb.WriteRune(' ')
	return sb.String()
}

func (el Heading) Content(text string) string {
	return text
}

func (el Heading) Closing() string {
	switch el.style {
	case printer.AnsiStyle:
		return "\033[22m"
	case printer.TviewStyle:
		return "[::B]"
	default:
		return ""
	}
}
