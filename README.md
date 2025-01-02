# reto

reto is a simple terminal web browser in [djot](https://djot.net/) markup

## build

``` sh
git clone https://github.com/arcxio/reto
cd reto
go build
```

## usage

reto expects either HTTP(S) URL or a local HTML file path as a single argument.

## controls

- j, down arrow: scroll down
- k, up arrow: scroll up
- g, home: scroll to top
- G, end: scroll to bottom
- ctrl+f, page down: scroll down by one page
- ctrl+b, page up: scroll up by one page
- H: move to previous page
- L: move to next page
- tab: select next link
- shift+tab: select previous link
- enter: follow selected link

## dependencies

- [github.com/rivo/tview](https://github.com/rivo/tview) - UI
- [github.com/gdamore/tcell/v2](https://github.com/gdamore/tcell) - terminal API
- [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) - HTML tokenizer

## license

reto is free software: you can redistribute it and/or modify it under the terms
of the GNU General Public License as published by the Free Software Foundation,
either version 3 of the License, or (at your option) any later version.

this program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. see the GNU General Public License for more details.
