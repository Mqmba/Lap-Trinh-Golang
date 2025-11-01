package student

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly/utils"
)

var studentList []Student

func addStudent() {
	// Implementation for adding a student
	var score []float64
	fmt.Println("-=-=-=-=-=- Thêm sinh viên mới -=-=-=-=-=-")
	id := utils.GetPositiveIntInput("Nhập mã số sinh viên: ")
	name := utils.ReadInput("Nhập họ và tên sinh viên: ")
	class := utils.ReadInput("Nhập lớp sinh viên: ")
	totalPoint := utils.GetPositiveIntInput("Nhập điểm sinh viên: ")
	for i := 1; i <= totalPoint; i++ {
		score = append(score, utils.GetPositiveFloatInput(fmt.Sprintf("- Nhập điểm môn thứ %d: ", i)))
	}

	student := Student{
		ID:    id,
		Name:  name,
		Class: class,
		Score: score,
	}
	studentList = append(studentList, student)
	fmt.Println("Đã thêm sinh viên thành công!")
}

func deleteStudent() {
	// Implementation for deleting a student
	var id int
	id = utils.GetPositiveIntInput("Nhập mã số sinh viên cần xóa: ")

	for i, student := range studentList {
		if student.ID == id {
			studentList = append(studentList[:i], studentList[i+1:]...)
			fmt.Println("Đã xóa sinh viên thành công!")
			break
		}
	}
	if len(studentList) == 0 || studentList[len(studentList)-1].ID != id {
		fmt.Println("Không tìm thấy sinh viên với mã số đã cho.")
	}
}

func showStudentList() {
	fmt.Println("-=-=-=-=-=- Danh sách sinh viên -=-=-=-=-=-")
	if len(studentList) == 0 {
		fmt.Println(" ==> Không có sinh viên nào trong danh sách.")
		return
	}
	for _, student := range studentList {
		fmt.Printf("%+v\n", student)
	}
}

func StudentMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("Quản lý sinh viên")
		fmt.Println("1️⃣. Thêm sinh viên")
		fmt.Println("2️⃣. Xóa sinh viên")
		fmt.Println("3️⃣. Hiển thị danh sách sinh viên")
		fmt.Println("4️⃣. ❌Thoát❌")

		choice := utils.GetPositiveIntInput("👉Vui lòng chọn một tùy chọn: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			showStudentList()
		case 4:
			fmt.Println("❌Thoát chương trình.❌")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục...")
	}
}
