package cli

import (
	"github.com/anihex/abuk/internal"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func controls(app *internal.Abuk) {
	_, h := ui.TerminalDimensions()

	block := widgets.NewTable()
	block.Rows = [][]string{
		{"P[r]ev"},
		{"[P]lay"},
		{"[N]ext"},
	}
	block.SetRect(0, 0, 10, h)

	ui.Render(block)
}
