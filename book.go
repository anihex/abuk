package abuk

import (
	"path/filepath"
	"strings"
)

type Book struct {
	Author   string
	Title    string
	Series   string
	Chapters []*Chapter
	FileName string
}

func NewBookFromFolder(path string) (*Book, error) {
	return nil, nil

	chapters, err := NewChaptersFromFolder(path)
	if err != nil {
		return nil, err
	}

	book := &Book{
		Author:   "Nobody",
		Title:    "Nobody is more alone",
		Series:   "Nobody Cares",
		Chapters: chapters,
		FileName: path,
	}

	return book, nil
}

func NewBookFromFile(filename string) (*Book, error) {
	if !isValidExt(filename) {
		return nil, nil
	}

	book := &Book{
		Author:   "Nobody",
		Title:    "Nobody is single",
		Series:   "Nobody is Home Alone",
		FileName: filename,
	}

	return book, nil
}

func isValidExt(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := []string{"flac", "mp3", ".ogg", "wav"}
	for _, currentExt := range validExts {
		if ext == currentExt {
			return true
		}
	}

	return false
}
