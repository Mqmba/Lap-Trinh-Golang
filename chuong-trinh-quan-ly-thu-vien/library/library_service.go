package library

import (
	"fmt"
	"time"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/utils"
)

func AddBook(lib *Library) error {
	id := utils.GenerateID()
	title := utils.GetNonEmptyString("Nhập tiêu đề: ")
	author := utils.GetNonEmptyString("Nhập tên tác giả: ")

	if err := lib.addBook(id, title, author); err != nil {
		return err
	}

	fmt.Printf("✅Thêm sách thành công! ID: %s\n", id)
	// fmt.Printf("%+v\n", lib.book)
	return nil
}

func ListBooks(lib *Library) error {
	// Lấy danh sách sách từ store
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Không có sách trong thư viện!")
		return nil
	}

	// Hiển thị danh sách sách
	for _, book := range books {
		status := "Còn"
		if book.IsBorrowed {
			status = "Đã mượn"
		}
		fmt.Printf("ID: %s, Tiêu đề: %s, Tác giả: %s, Trạng thái: %s\n", book.Id, book.Title, book.Author, status)
	}
	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateID()
	name := utils.GetNonEmptyString("Nhập tên người mượn: ")
	email := utils.GetNonEmptyString("Nhập email người mượn: ")
	if err := lib.addBorrower(id, name, email); err != nil {
		return err
	}
	fmt.Printf("✅Thêm người mượn thành công! ID: %s\n", id)
	return nil
}

func ListBorrowers(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Không có người mượn trong thư viện!")
		return nil
	}
	for _, borrower := range borrowers {
		fmt.Printf("ID: %s, Tên: %s, Email: %s\n", borrower.Id, borrower.Name, borrower.Email)
	}
	return nil
}

func BorrowBook(lib *Library) error {
	TransactionId := utils.GenerateID()
	bookId := utils.GetNonEmptyString("Nhập ID sách: ")
	borrowerId := utils.GetNonEmptyString("Nhập ID người mượn: ")

	if err := lib.Borrow(TransactionId, bookId, borrowerId); err != nil {
		return err
	}
	fmt.Printf("✅Mượn sách thành công! ID giao dịch: %s\n", TransactionId)
	return nil
}

func ListBorrowHistory(lib *Library) error {
	borrowerId := utils.GetNonEmptyString("Nhập ID người mượn: ")
	transactions := lib.ListBorrowHistoryByBorrower(borrowerId)

	if len(transactions) == 0 {
		fmt.Println("Không tìm thấy lịch sử mượn cho người mượn này hoặc người mượn không tồn tại!")
		return nil
	}

	fmt.Printf("Lịch sử mượn sách của người mượn ID: %s\n", borrowerId)
	for _, transaction := range transactions {
		returnDateStr := "Chưa trả"
		if !transaction.ReturnDate.IsZero() {
			returnDateStr = transaction.ReturnDate.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("ID giao dịch: %s, tên sách: %s, Ngày mượn: %s, Ngày trả: %s\n",
			transaction.TransactionId,
			lib.getBookTitle(transaction.BookId),
			transaction.BorrowDate.Format("2006-01-02 15:04:05"),
			returnDateStr)
	}
	return nil
}

func ReturnBook(lib *Library) error {
	transactionId := utils.GetNonEmptyString("Nhập ID giao dịch trả sách: ")
	transaction, exists := lib.transactions[transactionId]
	if !exists {
		return fmt.Errorf("Không tìm thấy giao dịch với ID: %s", transactionId)
	}
	if !transaction.ReturnDate.IsZero() {
		return fmt.Errorf("Sách đã được trả trước đó (Ngày trả: %s)", transaction.ReturnDate.Format("2006-01-02 15:04:05"))
	}

	// Cập nhật trạng thái sách
	book, exists := lib.books[transaction.BookId]
	if !exists {
		return fmt.Errorf("Không tìm thấy sách có ID: %s", transaction.BookId)
	}
	if !book.IsBorrowed {
		return fmt.Errorf("Sách đã được đánh dấu là chưa mượn.")
	}
	book.IsBorrowed = false
	lib.books[transaction.BookId] = book

	// Cập nhật ngày trả cho transaction
	transaction.ReturnDate = time.Now()
	lib.transactions[transactionId] = transaction

	fmt.Println("✅Trả sách thành công!")
	return nil
}

func SearchBook(lib *Library) error {
	query := utils.GetNonEmptyString("Nhập tên sách hoặc tác giả cần tìm: ")
	books := lib.SearchBookStore(query)
	if len(books) == 0 {
		fmt.Println("Không tìm thấy sách hoặc tác giả!")
		return nil
	}

	for _, book := range books {
		status := "Có sẵn"
		if book.IsBorrowed {
			status = "Đã cho mượn"
		}
		fmt.Printf("ID: %s, Tên: %s, Tác giả: %s, Trạng thái: %s\n", book.Id, book.Title, book.Author, status)
	}
	// foundBooks := []models.Book{}
	// for _, book := range lib.books {
	// 	if containsIgnoreCase(book.Title, query) {
	// 		foundBooks = append(foundBooks, book)
	// 	}
	// }
	// if len(foundBooks) == 0 {
	// 	fmt.Println("Không tìm thấy sách nào với tên này!")
	// 	return nil
	// }
	// fmt.Println("Kết quả tìm kiếm:")
	// for _, book := range foundBooks {
	// 	status := "Có sẵn"
	// 	if book.IsBorrowed {
	// 		status = "Đã cho mượn"
	// 	}
	// 	fmt.Printf("ID: %s, Tên: %s, Tác giả: %s, Trạng thái: %s\n", book.Id, book.Title, book.Author, status)
	// }
	return nil
}

// InitializeSampleData khởi tạo dữ liệu mẫu (2 sách và 2 người mượn)
func InitializeSampleData(lib *Library) {
	// Thêm 2 sách mẫu
	lib.addBook("book-001", "Đắc Nhân Tâm", "Dale Carnegie")
	lib.addBook("book-002", "Nhà Giả Kim", "Paulo Coelho")
	lib.addBook("book-003", "Tuổi Trẻ Đáng Giá Bao Nhiêu", "Rosie Nguyễn")
	lib.addBook("book-004", "Không Gia Đình", "Hector Malot")

	// Thêm 2 người mượn mẫu
	lib.addBorrower("borrower-001", "Nguyễn Văn A", "nguyenvana@example.com")
	lib.addBorrower("borrower-002", "Trần Thị B", "tranthib@example.com")
	lib.addBorrower("borrower-003", "Nguyễn Văn C", "nguyenvanc@example.com")

	// Thêm transaction
	lib.Borrow("transaction-001", "book-001", "borrower-001")
	lib.Borrow("transaction-003", "book-003", "borrower-001")
	lib.Borrow("transaction-002", "book-002", "borrower-002")
}
