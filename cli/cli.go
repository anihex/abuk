package cli

import (
	"log"

	"github.com/anihex/abuk/internal"
	ui "github.com/gizak/termui/v3"
)

func render(app *internal.Abuk) {
	chapters(app)
	books(app)
	controls(app)
	playing(app)
	status(app)
}

func CLI(app *internal.Abuk) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	render(app)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID { // event string/identifier
			case "q": // press 'q' or 'C-c' to quit
				return

			case "p":
				//go playTitle("")
			// case "<MouseLeft>":
			// 	payload := e.Payload.(ui.Mouse)
			// 	x, y := payload.X, payload.Y
			case "<Resize>":
				//payload := e.Payload.(ui.Resize)
				//width, height := payload.Width, payload.Height
				render(app)
			}
			switch e.Type {
			case ui.KeyboardEvent: // handle all key presses
				//eventID = e.ID // keypress string
			}
		}
	}
}
