package student

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly/utils"
)

var studentList []Student

func addStudent() {
	// Implementation for adding a student
	var score []float64
	var id int
	fmt.Println("-=-=-=-=-=- Thêm sinh viên mới -=-=-=-=-=-")
	for {
		id = utils.GetPositiveIntInput("Nhập mã số sinh viên: ")
		if !utils.CheckDuplicateID(id, studentList) {
			break
		}
		fmt.Println("❌Mã số sinh viên đã tồn tại. Vui lòng nhập mã số khác.❌")
	}
	name := utils.GetNonEmptyString("Nhập họ và tên sinh viên: ")
	class := utils.GetNonEmptyString("Nhập lớp sinh viên: ")
	totalPoint := utils.GetPositiveIntInput("Nhập điểm sinh viên: ")
	for i := 1; i <= totalPoint; i++ {
		var point float64
		for {
			point = utils.GetPositiveFloatInput(fmt.Sprintf("- Nhập điểm môn thứ %d: ", i))
			if point <= 10 {
				break
			}
			fmt.Println("❌Điểm số không hợp lệ [0-10]. Vui lòng nhập lại.❌")
		}
		score = append(score, point)
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

var studentSample = Student{
	ID:    1,
	Name:  "Nguyen Van A",
	Class: "CTK41",
	Score: []float64{8.5, 7.0, 9.0},
}

func init() {
	studentList = append(studentList, studentSample)
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

func updateStudent() {
	// Implementation for updating a student
	var id int
	id = utils.GetPositiveIntInput("Nhập mã số sinh viên cần sửa: ")
	student := searchStudentByID(id)
	if student != nil {
		fmt.Println("Nhập thông tin mới (Nhấn Enter để giữ nguyên giá trị hiện tại)")
		name := utils.GetOptionalString(fmt.Sprintf("Nhập họ và tên sinh viên mới(%s): ", student.Name), student.Name)
		class := utils.GetOptionalString(fmt.Sprintf("Nhập lớp sinh viên mới(%s): ", student.Class), student.Class)
		var score []float64
		for i, existingScore := range student.Score {
			point := utils.GetOptionalPositiveFloatInput(
				fmt.Sprintf("- Nhập điểm môn thứ %d (hiện tại: %.1f, Enter để giữ nguyên): ",
					i+1, existingScore), existingScore)
			score = append(score, point)
		}
		// Ask if user wants to add more scores
		fmt.Println("\nThêm điểm mới (nhập 0 để bỏ qua)")
		newPoints := utils.GetIntInput("Số điểm cần thêm: ")
		if newPoints > 0 {
			currentTotal := len(student.Score)
			for i := 1; i <= newPoints; i++ {
				var point float64
				for {
					point = utils.GetPositiveFloatInput(
						fmt.Sprintf("- Nhập điểm môn thứ %d: ", currentTotal+i))
					if point <= 10 {
						break
					}
					fmt.Println("❌Điểm số không hợp lệ [0-10]. Vui lòng nhập lại.❌")
				}
				score = append(score, point)
			}
		}
		// totalPoint := utils.GetPositiveIntInput("Nhập điểm sinh viên mới: ")
		// if totalPoint == 0 {
		// 	score = student.Score
		// } else {
		// 	for i := 1; i <= totalPoint; i++ {
		// 		score = append(score, utils.GetOptionalPositiveFloatInput(fmt.Sprintf("- Nhập điểm môn thứ %d: ", i), student.Score[i-1]))
		// 	}
		// }
		student.Name = name
		student.Class = class
		student.Score = score
		fmt.Println("Đã cập nhật thông tin sinh viên thành công!")
		fmt.Printf("%+v\n", *student)
	} else {
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
		fmt.Println(getInfo(student))
	}
	utils.ReadInput("Nhấn Enter để tiếp tục...")
}

func searchStudentByID(id int) *Student {
	for i, student := range studentList {
		if student.ID == id {
			return &studentList[i]
		}
	}
	return nil
}

func StudentMenu() {
	for {
		// utils.ClearScreen()
		fmt.Println("\n-=-=-=-=-=- Quản lý sinh viên -=-=-=-=-=-")
		fmt.Println("1️⃣. Thêm sinh viên")
		fmt.Println("2️⃣. Xóa sinh viên")
		fmt.Println("3️⃣. Sửa sinh viên")
		fmt.Println("4️⃣. Hiển thị danh sách sinh viên")
		fmt.Println("5️⃣. Tìm kiếm sinh viên theo mã số")
		fmt.Println("6️⃣. Quay lại menu chính")

		choice := utils.GetPositiveIntInput("👉Vui lòng chọn một tùy chọn: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			showStudentList()
		case 5:
			id := utils.GetPositiveIntInput("Nhập mã số sinh viên cần tìm: ")
			student := searchStudentByID(id)
			if student != nil {
				fmt.Println("Thông tin sinh viên:")
				fmt.Println(getInfo(*student))
			} else {
				fmt.Println("Không tìm thấy sinh viên với mã số đã cho.")
			}
		case 6:
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
	}
}
