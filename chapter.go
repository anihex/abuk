package abuk

import "time"

type Chapter struct {
	Title string
	File  string
	From  time.Time
	Until time.Time
}

func NewChaptersFromFolder(path string) ([]*Chapter, error) {
	return nil, nil
}
