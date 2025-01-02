package elements

import (
	"strconv"
	"strings"

	"github.com/arcxio/reto/internal/printer"
)

type Link struct {
	p        *printer.Printer
	url      string
	autolink bool
}

func (l *Link) Opening() string {
	for _, a := range l.p.Token().Attr {
		if a.Key == "href" {
			l.url = a.Val
			switch l.p.Style {
			case printer.AnsiStyle:
				var sb strings.Builder
				sb.WriteString("\033]8;;")
				sb.WriteString(l.url)
				sb.WriteString("\033[96;4m")
				return sb.String()
			case printer.TviewStyle:
				var sb strings.Builder
				sb.WriteString("[\"")
				sb.WriteString(strconv.Itoa(l.p.LinkCount()))
				sb.WriteString("\"]")
				sb.WriteString("[blue:::")
				sb.WriteString(l.url)
				sb.WriteRune(']')
				return sb.String()
			}
			break
		}
	}
	return ""
}

func (l *Link) Content(text string) string {
	if l.url != "" {
		var sb strings.Builder
		if text == strings.TrimPrefix(l.url, "mailto:") {
			l.autolink = true
			sb.WriteRune('<')
			sb.WriteString(text)
			sb.WriteRune('>')
		} else {
			sb.WriteRune('[')
			sb.WriteString(text)
		}
		l.p.PushLink(l.url, l.autolink)
		return sb.String()
	}
	return text
}

func (l Link) Closing() string {
	var sb strings.Builder
	if !l.autolink {
		if l.p.Style == printer.TviewStyle {
			sb.WriteString("[][")
		} else {
			sb.WriteString("][")
		}
		sb.WriteString(strconv.Itoa(l.p.LinkListCount()))
		sb.WriteRune(']')
	}
	switch l.p.Style {
	case printer.AnsiStyle:
		sb.WriteString("\033[0m\033]8;;\033\\")
	case printer.TviewStyle:
		sb.WriteString("[-:::-][\"\"]")
	}
	return sb.String()
}
