package library

import (
	"fmt"
	"strings"
	"time"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/models"
)

type Library struct {
	books        map[string]models.Book
	borrowers    map[string]models.Borrower
	transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

// addBook thêm sách vào thư viện
func (lib *Library) addBook(id string, title string, author string) error {
	// Check trùng ID(chắc chắn ko xảy ra vì dùng uuid)
	if _, exists := lib.books[id]; exists {
		return fmt.Errorf("Sách đã tồn tại với ID: %s", id)
	}
	lib.books[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return nil
}

// ListBooksStore lấy danh sách sách từ store
func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.books))
	for _, book := range lib.books {
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
	if _, exists := lib.books[bookId]; !exists {
		return fmt.Errorf("Sách không tồn tại với ID: %s", bookId)
	}
	if _, exists := lib.borrowers[borrowerId]; !exists {
		return fmt.Errorf("Người mượn không tồn tại với ID: %s", borrowerId)
	}
	if _, exists := lib.transactions[id]; exists {
		return fmt.Errorf("Giao dịch đã tồn tại với ID: %s", id)
	}

	// Lấy sách từ store
	book, bookExists := lib.books[bookId]
	if !bookExists {
		return fmt.Errorf("Sách không tồn tại với ID: %s", bookId)
	}

	// Nếu sách đã được mượn thì trả lỗi
	if book.IsBorrowed {
		return fmt.Errorf("Sách %s đã được mượn", book.Title)
	}

	// Cập nhật trạng thái sách
	book.IsBorrowed = true
	lib.books[bookId] = book

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

// ListBorrowHistoryStore lấy danh sách lịch sử mượn từ store
func (lib *Library) ListBorrowHistoryByBorrower(borrowerId string) []models.Transaction {
	// Kiểm tra người mượn có tồn tại không
	if _, exists := lib.borrowers[borrowerId]; !exists {
		return nil
	}

	// Lọc các giao dịch theo borrowerId
	transactions := make([]models.Transaction, 0)
	for _, transaction := range lib.transactions {
		if transaction.BorrowerId == borrowerId {
			transactions = append(transactions, transaction)
		}
	}
	return transactions
}

func (lib *Library) getBookTitle(bookId string) string {
	book := lib.books[bookId]
	return book.Title
}

func (lib *Library) SearchBookStore(query string) []models.Book {
	query = strings.ToLower(query)
	var results []models.Book
	for _, book := range lib.books {
		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Author), query) {
			results = append(results, book)
		}
	}
	return results
}

// Hàm tiện ích so sánh không phân biệt hoa thường
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
