package main

import (
	"io"
	"strings"
	"testing"

	"github.com/arcxio/reto/internal/printer"
)

func assert(t *testing.T, style printer.Style, input, expected string) {
	t.Helper()
	var sb strings.Builder
	Print(io.NopCloser(strings.NewReader(input)), &sb, printer.NewPrinter(style))
	if actual := sb.String(); expected != actual {
		t.Error("\nexpected: " + expected + "\nactual  : " + actual)
	}
}

func TestHeading(t *testing.T) {
	t.Run("level 1", func(t *testing.T) {
		input := "<h1>Level 1 heading</h1>"
		assert(t, printer.NoStyle, input, "# Level 1 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m# Level 1 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]# Level 1 heading[::B]")
	})

	t.Run("level 2", func(t *testing.T) {
		input := "<h2>Level 2 heading</h2>"
		assert(t, printer.NoStyle, input, "## Level 2 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m## Level 2 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]## Level 2 heading[::B]")
	})
	t.Run("level 3", func(t *testing.T) {
		input := "<h3>Level 3 heading</h3>"
		assert(t, printer.NoStyle, input, "### Level 3 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m### Level 3 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]### Level 3 heading[::B]")
	})
	t.Run("level 4", func(t *testing.T) {
		input := "<h4>Level 4 heading</h4>"
		assert(t, printer.NoStyle, input, "#### Level 4 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m#### Level 4 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]#### Level 4 heading[::B]")
	})
	t.Run("level 5", func(t *testing.T) {
		input := "<h5>Level 5 heading</h5>"
		assert(t, printer.NoStyle, input, "##### Level 5 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m##### Level 5 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]##### Level 5 heading[::B]")
	})
	t.Run("level 6", func(t *testing.T) {
		input := "<h6>Level 6 heading</h6>"
		assert(t, printer.NoStyle, input, "###### Level 6 heading")
		assert(t, printer.AnsiStyle, input, "\033[1m###### Level 6 heading\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]###### Level 6 heading[::B]")
	})
}

func TestLink(t *testing.T) {
	t.Run("link", func(t *testing.T) {
		input := `<a href="url">link</a>`
		assert(t, printer.NoStyle, input, "[link][1]\n\n[1]: url")
		assert(t, printer.AnsiStyle, input, "\033]8;;url\033[96;4m[link][1]\033[0m\033]8;;\033\\\n\n[1]: \033]8;;url\033[96;4murl\033[0m\033]8;;\033\\")
		assert(t, printer.TviewStyle, input, "[\"0\"][blue:::url][link[][1][-:::-][\"\"]\n\n[1]: [blue:::url]url[-:::-]")
	})
	t.Run("link emphasised", func(t *testing.T) {
		input := `<a href="url">link <em>emphasised</em></a>`
		assert(t, printer.NoStyle, input, "[link _emphasised_][1]\n\n[1]: url")
		assert(t, printer.AnsiStyle, input, "\033]8;;url\033[96;4m[link \033[3m_emphasised_\033[24m][1]\033[0m\033]8;;\033\\\n\n[1]: \033]8;;url\033[96;4murl\033[0m\033]8;;\033\\")
		assert(t, printer.TviewStyle, input, "[\"0\"][blue:::url][link [::u]_emphasised_[::U][][1][-:::-][\"\"]\n\n[1]: [blue:::url]url[-:::-]")
	})
	t.Run("autolink", func(t *testing.T) {
		input := `<a href="https://example.org">https://example.org</a>`
		assert(t, printer.NoStyle, input, "<https://example.org>")
		assert(t, printer.AnsiStyle, input, "\033]8;;https://example.org\033[96;4m<https://example.org>\033[0m\033]8;;\033\\")
		assert(t, printer.TviewStyle, input, "[\"0\"][blue:::https://example.org]<https://example.org>[-:::-][\"\"]")
	})
	t.Run("email", func(t *testing.T) {
		input := `<a href="mailto:me@example.com">me@example.com</a>`
		assert(t, printer.NoStyle, input, "<me@example.com>")
		assert(t, printer.AnsiStyle, input, "\033]8;;mailto:me@example.com\033[96;4m<me@example.com>\033[0m\033]8;;\033\\")
		assert(t, printer.TviewStyle, input, "[\"0\"][blue:::mailto:me@example.com]<me@example.com>[-:::-][\"\"]")
	})
}

func TestEmphasis(t *testing.T) {
	t.Run("i", func(t *testing.T) {
		input := "<i>foo bar</i>"
		assert(t, printer.NoStyle, input, "_foo bar_")
		assert(t, printer.AnsiStyle, input, "\033[3m_foo bar_\033[24m")
		assert(t, printer.TviewStyle, input, "[::u]_foo bar_[::U]")
	})
	t.Run("u", func(t *testing.T) {
		input := "<u>foo bar</u>"
		assert(t, printer.NoStyle, input, "_foo bar_")
		assert(t, printer.AnsiStyle, input, "\033[3m_foo bar_\033[24m")
		assert(t, printer.TviewStyle, input, "[::u]_foo bar_[::U]")
	})
	t.Run("em", func(t *testing.T) {
		input := "<em>foo bar</em>"
		assert(t, printer.NoStyle, input, "_foo bar_")
		assert(t, printer.AnsiStyle, input, "\033[3m_foo bar_\033[24m")
		assert(t, printer.TviewStyle, input, "[::u]_foo bar_[::U]")
	})
	t.Run("b", func(t *testing.T) {
		input := "<b>foo bar</b>"
		assert(t, printer.NoStyle, input, "*foo bar*")
		assert(t, printer.AnsiStyle, input, "\033[1m*foo bar*\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]*foo bar*[::B]")
	})
	t.Run("strong", func(t *testing.T) {
		input := "<strong>foo bar</strong>"
		assert(t, printer.NoStyle, input, "*foo bar*")
		assert(t, printer.AnsiStyle, input, "\033[1m*foo bar*\033[22m")
		assert(t, printer.TviewStyle, input, "[::b]*foo bar*[::B]")
	})
}

func TestDefinitionList(t *testing.T) {
	assert(t, printer.NoStyle,
		"<dl><dt>apple</dt><dd><p>red fruit</p></dd><dt>banana</dt><dd><p>yellow fruit</p></dd></dl>",
		": apple\n\n  red fruit\n\n: banana\n\n  yellow fruit",
	)
}

func TestList(t *testing.T) {
	assert(t, printer.NoStyle,
		"<ul><li>a<ul><li>b</li><li>c</li></ul></li><li>d</li></ul>",
		"- a\n\n  - b\n\n  - c\n\n- d",
	)
}

func TestVerbatim(t *testing.T) {
	assert(t, printer.NoStyle,
		"<p>This is <mark>highlighted text</mark>.</p>",
		"This is {=highlighted text=}.",
	)
}

func TestSuperscriptAndSupscript(t *testing.T) {
	assert(t, printer.NoStyle,
		"<p>H<sub>2</sub>O and djot<sup>TM</sup></p>",
		"H~2~O and djot^TM^",
	)
}

func TestInsertAndDelete(t *testing.T) {
	assert(t, printer.NoStyle,
		"<p>My boss is <del>mean</del><ins>nice</ins>.</p>",
		"My boss is {-mean-}{+nice+}.",
	)
}

func TestNewline(t *testing.T) {
	assert(t, printer.NoStyle,
		"<p>This is a soft\nbreak and this is a hard<br>\nbreak.</p>",
		"This is a soft break and this is a hard\\\nbreak.",
	)
}

func TestThematicBreak(t *testing.T) {
	assert(t, printer.NoStyle,
		"<p>Then they went to sleep.</p><hr><p>When they woke up, &hellip;</p>",
		"Then they went to sleep.\n\n****\n\nWhen they woke up, …",
	)
}
