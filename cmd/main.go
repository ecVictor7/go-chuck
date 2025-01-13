package main

import (
	"github.com/rivo/tview"
)

var (
	app *tview.Application
)

func main() {
	app = tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("Hello world from Tview")

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
