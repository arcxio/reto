package elements

type insert struct{}

func (_ insert) Opening() string {
	return "{+"
}

func (_ insert) Content(text string) string {
	return text
}

func (_ insert) Closing() string {
	return "+}"
}
