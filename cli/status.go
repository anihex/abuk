package cli

import (
	"github.com/anihex/abuk/internal"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func status(app *internal.Abuk) {
	w, _ := ui.TerminalDimensions()
	block := widgets.NewParagraph()
	block.SetRect(w-10, 0, w, 6)
	block.Text = `00:00:00
	00:00:00
	00:00:00
	00:00:00`

	ui.Render(block)
}
