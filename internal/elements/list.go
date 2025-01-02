package elements

import (
	"strings"

	"golang.org/x/net/html/atom"

	"github.com/arcxio/reto/internal/printer"
)

type ListItem struct {
	p *printer.Printer
}

func (el ListItem) Opening() string {
	var sb strings.Builder
	for i := 1; i < el.p.AtomDepth(atom.Ul); i++ {
		sb.WriteString("  ")
	}
	sb.WriteString("- ")
	return sb.String()
}

func (el ListItem) Content(text string) string {
	return text
}

func (el ListItem) Closing() string {
	return ""
}
