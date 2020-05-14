package cli

import (
	"github.com/anihex/abuk/internal"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func playing(app *internal.Abuk) {
	w, _ := ui.TerminalDimensions()
	block := widgets.NewParagraph()
	block.SetRect(10, 0, w-10, 6)
	block.Title = " Now Playing "
	block.Text = `Author : Joan K. Rowling
	Book   : Harry Potter und der Stein der Weisen - Harry Potter 1 - Gelesen von Rufus Beck
	Chapter: 01 - Ein Junge überlebt`

	ui.Render(block)
}
