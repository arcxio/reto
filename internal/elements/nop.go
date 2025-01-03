package elements

type Nop struct{}

func (_ Nop) Opening() string {
	return ""
}

func (_ Nop) Content(text string) string {
	return text
}

func (_ Nop) Closing() string {
	return ""
}
