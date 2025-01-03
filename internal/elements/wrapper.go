package elements

type Wrapper struct {
	opening string
	closing string
}

func (el Wrapper) Opening() string {
	return el.opening
}

func (_ Wrapper) Content(text string) string {
	return text
}

func (el Wrapper) Closing() string {
	return el.closing
}
