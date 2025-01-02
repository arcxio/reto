package main

import (
	"fmt"
	"os"
)

const Version = "0.1"

func die(cause string) {
	fmt.Fprintln(os.Stderr, os.Args[0]+": "+cause)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		die("usage: reto [page]")
	}

	pages := NewPages()
	if err := pages.Open(os.Args[1]); err != nil {
		die(err.Error())
	}

	app := NewApplication(pages)
	if err := app.Run(); err != nil {
		app.Stop()
		die(err.Error())
	}
}
