package library

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/models"
	"mamba.com/chuong-trinh-quan-ly-thu-vien/utils"
)

func AddBook(lib *Library) error {
	id := utils.GenerateID()
	title := utils.GetNonEmptyString("Nhập tiêu đề: ")
	author := utils.GetNonEmptyString("Nhập tên tác giả: ")

	lib.book[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	fmt.Println("Thêm sách thành công!")
	fmt.Printf("%+v\n", lib.book)
	return nil
}

func ListBooks() error {
	return nil
}

func AddBorrower() error {
	return nil
}

func ListBorrowers() error {
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
