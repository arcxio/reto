package printer

type Style byte

const (
	NoStyle Style = iota
	AnsiStyle
	TviewStyle
)
