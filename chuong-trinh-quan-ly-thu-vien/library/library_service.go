package library

import (
	"fmt"

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

func BorrowBook() error {
	return nil
}

func ListBorrowHistory() error {
	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBook() error {
	return nil
}
