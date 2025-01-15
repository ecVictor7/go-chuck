package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const refreshInterval = 10 * time.Second
const url = "https://api.chucknorris.io/jokes/random"

var (
	app      *tview.Application
	textView *tview.TextView
)

type Payload struct {
	Value string
}

func getAndDrawJoke() {
	//fetch chuck norris joke from the web
	result, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	payloadBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	payload := &Payload{}
	err = json.Unmarshal(payloadBytes, payload)
	if err != nil {
		panic(err)
	}

	//update our UI with the joke
	textView.Clear()
	fmt.Fprintln(textView, payload.Value)
	timeStr := fmt.Sprintf("\n\n[gray]%s", time.Now().Format(time.RFC1123))
	fmt.Fprintln(textView, timeStr)
}

func refreshJoke() {
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			getAndDrawJoke()
			app.Draw()
		}
	}
}
func main() {
	app = tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLime)

	textView.SetBorderPadding(1, 0, 0, 0)

	getAndDrawJoke()
	go refreshJoke()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
