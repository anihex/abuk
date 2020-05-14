package internal

import (
	"log"

	"github.com/anihex/abuk"
)

type Abuk struct {
	Playlist abuk.Playlist
}

func NewAbuk(config Config) *Abuk {
	playlist, err := abuk.NewPlaylist(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The Folder contains %d books", len(playlist.Books))
	for index, book := range playlist.Books {
		log.Printf("Book %d: %s (%s)", index, book.Title, book.FileName)
	}

	return &Abuk{
		Playlist: *playlist,
	}
}
