package teacher

import (
	"fmt"

	"mamba.com/chuong-trinh-quan-ly/utils"
)

var teacherList []Teacher

func addTeacher() {
	// Implementation for adding a teacher
	var id int
	fmt.Println("-=-=-=-=-=- Thêm giảng viên mới -=-=-=-=-=-")
	for {
		id = utils.GetPositiveIntInput("Nhập mã số giảng viên: ")
		if !utils.CheckDuplicateID(id, teacherList) {
			break
		}
		fmt.Println("❌Mã số giảng viên đã tồn tại. Vui lòng nhập mã số khác.❌")
	}
	name := utils.ReadInput("Nhập họ và tên giảng viên: ")
	subjects := utils.ReadInput("Nhập môn giảng dạy: ")
	salary := utils.GetPositiveFloatInput("Nhập lương cơ bản: ")
	bonus := utils.GetPositiveFloatInput("Nhập tiền thưởng: ")

	teacher := Teacher{
		ID:       id,
		Name:     name,
		subjects: subjects,
		salary:   salary,
		bonus:    bonus,
	}
	teacherList = append(teacherList, teacher)
	fmt.Println("Đã thêm giảng viên thành công!")
}

func deleteTeacher() {
	// Implementation for deleting a teacher
	var id int
	id = utils.GetPositiveIntInput("Nhập mã số giảng viên cần xóa: ")
	for i, teacher := range teacherList {
		if teacher.ID == id {
			teacherList = append(teacherList[:i], teacherList[i+1:]...)
			fmt.Println("Đã xóa giảng viên thành công!")
			break
		}
	}
	if len(teacherList) == 0 || teacherList[len(teacherList)-1].ID != id {
		fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
	}
}

func showTeacherList() {
	fmt.Println("-=-=-=-=-=- Danh sách giảng viên -=-=-=-=-=-")
	if len(teacherList) == 0 {
		fmt.Println(" ==> Không có giảng viên nào trong danh sách.")
		return
	}
	for _, teacher := range teacherList {
		fmt.Println(getInfo(teacher))
	}
	utils.ReadInput("Nhấn Enter để tiếp tục...")
}

func searchTeacherByID(id int) *Teacher {
	for i, teacher := range teacherList {
		if teacher.ID == id {
			return &teacherList[i]
		}
	}
	return nil
}

func TeacherMenu() {
	for {
		// utils.ClearScreen()
		fmt.Println("\n-=-=-=-=-=- Quản lý giảng viên -=-=-=-=-=-")
		fmt.Println("1. Thêm giảng viên")
		fmt.Println("2. Xóa giảng viên")
		fmt.Println("3. Hiển thị danh sách giảng viên")
		fmt.Println("4. Tìm kiếm giảng viên theo mã số")
		fmt.Println("5. Quay lại menu chính")
		choice := utils.GetPositiveIntInput("Chọn một tùy chọn (1-5): ")
		switch choice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			showTeacherList()
		case 4:
			id := utils.GetPositiveIntInput("Nhập mã số giảng viên cần tìm: ")
			teacher := searchTeacherByID(id)
			if teacher != nil {
				fmt.Println("Thông tin giảng viên:")
				fmt.Println(getInfo(*teacher))
			} else {
				fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
			}
		case 5:
			return
		default:
			fmt.Println("Tùy chọn không hợp lệ. Vui lòng chọn lại.")
		}
	}
}
