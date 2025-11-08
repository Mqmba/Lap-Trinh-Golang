package library

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/models"
)

type Library struct {
	book      map[string]models.Book
	borrowers map[string]models.Borrower
}

func NewLibrary() *Library {
	return &Library{
		book:      make(map[string]models.Book),
		borrowers: make(map[string]models.Borrower),
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

// ListBooksStore lấy danh sách sách từ store
func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.book))
	for _, book := range lib.book {
		books = append(books, book)
	}
	return books
}

// addBorrower thêm người mượn vào thư viện
func (lib *Library) addBorrower(id string, name string, email string) error {
	if _, exists := lib.borrowers[id]; exists {
		return fmt.Errorf("Người mượn đã tồn tại với ID: %s", id)
	}
	lib.borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}
	return nil
}

// ListBorrowersStore lấy danh sách người mượn từ store
func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrowers))
	for _, borrower := range lib.borrowers {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}
