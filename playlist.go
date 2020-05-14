package abuk

import "io/ioutil"

type Playlist struct {
	Books []*Book
}

func NewPlaylist(path string) (*Playlist, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	playlist := &Playlist{
		Books: make([]*Book, 0),
	}

	for _, file := range files {
		if file.IsDir() {
			book, err := NewBookFromFolder(file.Name())
			if err != nil {
				return nil, err
			}

			if book != nil {
				playlist.Books = append(playlist.Books, book)
			}
		} else {
			book, err := NewBookFromFile(file.Name())
			if err != nil {
				return nil, err
			}

			if book != nil {
				playlist.Books = append(playlist.Books, book)
			}
		}
	}

	return playlist, nil
}
