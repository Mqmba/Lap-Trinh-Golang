package library

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/models"
)

type Library struct {
	book map[string]models.Book
}

func NewLibrary() *Library {
	return &Library{
		book: make(map[string]models.Book),
	}
}

// addBook thêm sách vào thư viện
func (lib *Library) addBook(id string, title string, author string) error {
	// Check trùng ID(chắc chắn ko xảy ra vì dùng uuid)
	if _, exists := lib.book[id]; exists {
		return fmt.Errorf("Sách đã tồn tại với ID: %s", id)
	}
	lib.book[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.book))
	for _, book := range lib.book {
		books = append(books, book)
	}
	return books
}
