package main

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly-thu-vien/library"
	"mamba.com/chuong-trinh-quan-ly-thu-vien/utils"
)

func main() {

	lib := library.NewLibrary()

	for {
		utils.ClearScreen()
		fmt.Println("CHƯƠNG TRÌNH QUẢN LÝ THƯ VIỆN")
		fmt.Println("1️⃣. Thêm sách")
		fmt.Println("2️⃣. Xem danh sách sách")
		fmt.Println("3️⃣. Thêm người mượn")
		fmt.Println("4️⃣. Xem danh sách người mượn")
		fmt.Println("5️⃣. Mượn sách")
		fmt.Println("6️⃣. Xem lịch sử mượn")
		fmt.Println("7️⃣. Trả sách")
		fmt.Println("8️⃣. Tìm kiếm sách")
		fmt.Println("9️⃣. Thoát")

		choice := utils.GetPositiveIntInput("👉Vui lòng chọn một tùy chọn: ")

		utils.ClearScreen()

		switch choice {
		case 1:
			fmt.Println("-=-=-=-=- Thêm sách -=-=-=-=-")
			if err := library.AddBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm sách: %+v \n", err)
			}
		case 2:
			fmt.Println("-=-=-=-=- Xem danh sách sách -=-=-=-=-")
			if err := library.ListBooks(); err != nil {
				fmt.Printf("❌ Lỗi khi xem danh sách sách: %+v \n", err)
			}
		case 3:
			fmt.Println("-=-=-=-=- Thêm người mượn -=-=-=-=-")
			if err := library.AddBorrower(); err != nil {
				fmt.Printf("❌ Lỗi khi thêm người mượn: %+v \n", err)
			}
		case 4:
			fmt.Println("-=-=-=-=- Xem danh sách người mượn -=-=-=-=-")
			if err := library.ListBorrowers(); err != nil {
				fmt.Printf("❌ Lỗi khi xem danh sách người mượn: %+v \n", err)
			}
		case 5:
			fmt.Println("-=-=-=-=- Mượn sách -=-=-=-=-")
			if err := library.BorrowBook(); err != nil {
				fmt.Printf("❌ Lỗi khi mượn sách: %+v \n", err)
			}
		case 6:
			fmt.Println("-=-=-=-=- Xem lịch sử mượn -=-=-=-=-")
			if err := library.ListBorrowHistory(); err != nil {
				fmt.Printf("❌ Lỗi khi xem lịch sử mượn: %+v \n", err)
			}
		case 7:
			fmt.Println("-=-=-=-=- Trả sách -=-=-=-=-")
			if err := library.ReturnBook(); err != nil {
				fmt.Printf("❌ Lỗi khi trả sách: %+v \n", err)
			}
		case 8:
			fmt.Println("-=-=-=-=- Tìm kiếm sách -=-=-=-=-")
			if err := library.SearchBook(); err != nil {
				fmt.Printf("❌ Lỗi khi tìm kiếm sách: %+v \n", err)
			}
		case 9:
			fmt.Println("❌Thoát chương trình.❌")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục...")
	}
}
