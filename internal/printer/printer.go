package printer

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type link struct {
	url      string
	autolink bool
}

type Printer struct {
	Style Style

	blockIndented      bool
	definitionIndented bool
	size               int
	linkListCount      int
	links              []link
	formatters         []string
	tokenStack         []html.Token
	atomDepths         map[atom.Atom]int
}

func (p *Printer) Size() int {
	return p.size
}

func (p *Printer) PushToken(token html.Token) {
	p.tokenStack = append(p.tokenStack, token)
	p.atomDepths[token.DataAtom]++
}

func (p *Printer) PopToken() (token *html.Token) {
	if token = p.Token(); token != nil {
		p.tokenStack = p.tokenStack[:len(p.tokenStack)-1]
		p.atomDepths[token.DataAtom]--
	}
	return token
}

func (p *Printer) Token() *html.Token {
	if len(p.tokenStack) > 0 {
		return &p.tokenStack[len(p.tokenStack)-1]
	}
	return nil
}

func (p *Printer) AtomDepth(a atom.Atom) int {
	return p.atomDepths[a]
}

func (p *Printer) InAtom(atoms ...atom.Atom) bool {
	for _, a := range atoms {
		if p.AtomDepth(a) > 0 {
			return true
		}
	}
	return false
}

func Link(url, text string, style Style) string {
	switch style {
	case AnsiStyle:
		var sb strings.Builder
		sb.WriteString("\033]8;;")
		sb.WriteString(url)
		sb.WriteString("\033[96;4m")
		sb.WriteString(text)
		sb.WriteString("\033[0m\033]8;;\033\\")
		return sb.String()
	case TviewStyle:
		var sb strings.Builder
		sb.WriteString("[blue:::")
		sb.WriteString(url)
		sb.WriteRune(']')
		sb.WriteString(text)
		sb.WriteString("[-:::-]")
		return sb.String()
	}
	return text
}

func (p *Printer) print(w io.Writer, format string, a ...any) (n int, err error) {
	if len(a) > 0 {
		n, err = fmt.Fprintf(w, format, a...)
	} else {
		n, err = fmt.Fprint(w, format)
	}
	p.size = p.size + n
	return
}

func (p *Printer) LinkCount() int {
	return len(p.links)
}

func (p *Printer) LinkListCount() int {
	return p.linkListCount
}

func (p *Printer) LinkUrl(i int) *string {
	if p.LinkCount() > i {
		return &p.links[i].url
	}
	return nil
}

func (p *Printer) Print(w io.Writer, a string) error {
	if a != "" {
		if n, err := p.print(w, a); err != nil {
			return err
		} else if n > 0 {
			p.blockIndented = false
			p.definitionIndented = false
		}
	}
	return nil
}

func SetTitle(title string) (int, error) {
	if title == "" {
		return fmt.Print("\033]0;reto\007")
	} else {
		return fmt.Print("\033]0;" + title + " - reto\007")
	}
}

func (p *Printer) SetTitle(title string) {
	if p.Style != NoStyle {
		SetTitle(title)
	}
}

func (p *Printer) IndentBlock(w io.Writer) error {
	if p.size > 0 && !p.blockIndented {
		if _, err := p.print(w, "\n\n"); err != nil {
			return err
		}
		p.blockIndented = true
	}
	return nil
}

func (p *Printer) IndentDefinition(w io.Writer) error {
	if !p.definitionIndented {
		if _, err := p.print(w, "  "); err != nil {
			return err
		}
		p.definitionIndented = true
	}
	return nil
}

func (p *Printer) PushLink(url string, autolink bool) {
	p.links = append(p.links, link{url: url, autolink: autolink})
	if !autolink {
		p.linkListCount++
	}
}

func (p *Printer) PushFormatter(formatter string) {
	p.formatters = append(p.formatters, formatter)
}

func (p *Printer) PopFormatter() {
	if len(p.formatters) > 0 {
		p.formatters = p.formatters[:len(p.formatters)-1]
	}
}

func (p *Printer) Formatters() []string {
	return p.formatters
}

func (p *Printer) PrintLinks(w io.Writer) error {
	if p.linkListCount > 0 {
		if _, err := p.print(w, "\n"); err != nil {
			return err
		}
		i := 1
		for _, l := range p.links {
			if !l.autolink {
				if _, err := p.print(w, "\n[%d]: %s", i, Link(l.url, l.url, p.Style)); err != nil {
					return err
				}
				i++
			}
		}
	}
	return nil
}

func NewPrinter(style Style) *Printer {
	p := Printer{Style: style}
	p.atomDepths = make(map[atom.Atom]int)
	return &p
}
