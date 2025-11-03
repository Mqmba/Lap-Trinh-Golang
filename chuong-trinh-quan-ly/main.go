package main

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly/student"
	"mamba.com/chuong-trinh-quan-ly/teacher"
	"mamba.com/chuong-trinh-quan-ly/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("Chương trình quản lý")
		fmt.Println("1️⃣. Quản lý sinh viên")
		fmt.Println("2️⃣. Quản lý giảng viên")
		fmt.Println("3️⃣. ❌Thoát❌")

		choice := utils.GetPositiveIntInput("👉Vui lòng chọn một tùy chọn: ")

		switch choice {
		case 1:
			student.StudentMenu()
		case 2:
			teacher.TeacherMenu()
		case 3:
			fmt.Println("❌Thoát chương trình.❌")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục...")
	}
}
