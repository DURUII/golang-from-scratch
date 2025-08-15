package model

import "fmt"

type Library struct {
	Name      string
	Books     []*Book
	Magazines []*Magazine
}

func (lib *Library) AddBook(book Book) bool {
	lib.Books = append(lib.Books, &book)
	return true
}

func (lib *Library) AddMagazine(magazine Magazine) bool {
	lib.Magazines = append(lib.Magazines, &magazine)
	return true
}

func (lib *Library) GetBookByID(id int) *Book {
	for _, book := range lib.Books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func (lib *Library) GetMagazineByID(id int) *Magazine {
	for _, magazine := range lib.Magazines {
		if magazine.ID == id {
			return magazine
		}
	}
	return nil
}

func (lib *Library) ListAllAvailableBooks() []Book {
	var ret []Book
	for _, book := range lib.Books {
		if book.IsAvailable {
			ret = append(ret, *book)
		}
	}
	return ret
}

func (lib *Library) ListAllAvailableMagazines() []Magazine {
	var ret []Magazine
	for _, magazine := range lib.Magazines {
		if magazine.IsAvailable {
			ret = append(ret, *magazine)
		}
	}
	return ret
}

func (lib *Library) PrintAllAvailableItems() {
	for _, book := range lib.ListAllAvailableBooks() {
		fmt.Println(book.GetInfo())
	}
	for _, magazine := range lib.ListAllAvailableMagazines() {
		fmt.Println(magazine.GetInfo())
	}
}
