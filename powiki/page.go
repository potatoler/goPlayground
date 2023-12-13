package powiki

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) Save() (err error) {
	return os.WriteFile(page.Title, page.Body, 0666)
}

func (page *Page) Load(title string) (err error) {
	page.Title = title
	page.Body, err = os.ReadFile(page.Title)
	return err
}
