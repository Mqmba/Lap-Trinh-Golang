package library

import (
	"fmt"
	"time"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/models"
)

type Library struct {
	book         map[string]models.Book
	borrowers    map[string]models.Borrower
	transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		book:         make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
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

// Borrow mượn sách
func (lib *Library) Borrow(id string, bookId string, borrowerId string) error {
	// Kiểm tra sách và người mượn có tồn tại không và giao dịch đã tồn tại chưa
	if _, exists := lib.book[bookId]; !exists {
		return fmt.Errorf("Sách không tồn tại với ID: %s", bookId)
	}
	if _, exists := lib.borrowers[borrowerId]; !exists {
		return fmt.Errorf("Người mượn không tồn tại với ID: %s", borrowerId)
	}
	if _, exists := lib.transactions[id]; exists {
		return fmt.Errorf("Giao dịch đã tồn tại với ID: %s", id)
	}

	// Lấy sách từ store
	book, bookExists := lib.book[bookId]
	if !bookExists {
		return fmt.Errorf("Sách không tồn tại với ID: %s", bookId)
	}

	// Nếu sách đã được mượn thì trả lỗi
	if book.IsBorrowed {
		return fmt.Errorf("Sách %s đã được mượn", book.Title)
	}

	// Cập nhật trạng thái sách
	book.IsBorrowed = true
	lib.book[bookId] = book

	// Tạo giao dịch mượn sách
	lib.transactions[id] = models.Transaction{
		TransactionId: id,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowDate:    time.Now(),
		ReturnDate:    time.Time{}, // Trả sách chưa có ngày trả
	}
	return nil
}
