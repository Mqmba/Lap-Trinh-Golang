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
	name := utils.GetNonEmptyString("Nhập họ và tên giảng viên: ")
	subjects := utils.GetNonEmptyString("Nhập môn giảng dạy: ")
	salary := utils.GetPositiveFloatInput("Nhập lương cơ bản: ")
	bonus := utils.GetPositiveFloatInput("Nhập tiền thưởng: ")

	teacher := Teacher{
		ID:       id,
		Name:     name,
		Subjects: subjects,
		Salary:   salary,
		Bonus:    bonus,
	}
	teacherList = append(teacherList, teacher)
	fmt.Println("Đã thêm giảng viên thành công!")
	utils.ReadInput("Nhấn Enter để tiếp tục")
}

// Teacher sample
var teacherSample = Teacher{
	ID:       1,
	Name:     "Nguyen Van C",
	Subjects: "CNTT",
	Salary:   10000000,
	Bonus:    200000,
}

func init() {
	teacherList = append(teacherList, teacherSample)
}

func deleteTeacher() {
	// Implementation for deleting a teacher
	if len(teacherList) == 0 {
		fmt.Println("Danh sách giảng viên trống.")
		utils.ReadInput("Nhấn Enter để tiếp tục")
		return
	}
	var id int
	id = utils.GetPositiveIntInput("Nhập mã số giảng viên cần xóa: ")
	for i, teacher := range teacherList {
		if teacher.ID == id {
			teacherList = append(teacherList[:i], teacherList[i+1:]...)
			fmt.Println("Đã xóa giảng viên thành công!")
			utils.ReadInput("Nhấn Enter để tiếp tục")
			return
		}
	}
	fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
	utils.ReadInput("Nhấn Enter để tiếp tục")
}

func updateTeacher() {
	// Implementation for updating a student
	var id int
	id = utils.GetPositiveIntInput("Nhập mã số giảng viên cần sửa: ")
	teacher := searchTeacherByID(id)
	if teacher != nil {
		fmt.Println("Nhập thông tin mới (Nhấn Enter để giữ nguyên giá trị hiện tại)")
		name := utils.GetOptionalString(fmt.Sprintf("Nhập họ và tên giảng viên mới(%s): ", teacher.Name), teacher.Name)
		subjects := utils.GetOptionalString(fmt.Sprintf("Nhập môn giảng dạy mới(%s): ", teacher.Subjects), teacher.Subjects)
		salary := utils.GetOptionalPositiveFloatInput(fmt.Sprintf("Nhập lương cơ bản(%.2f): ", teacher.Salary), teacher.Salary)
		bonus := utils.GetOptionalPositiveFloatInput(fmt.Sprintf("Nhập tiền thưởng(%.2f): ", teacher.Bonus), teacher.Bonus)
		teacher.Name = name
		teacher.Subjects = subjects
		teacher.Salary = salary
		teacher.Bonus = bonus
		fmt.Println("Đã cập nhật thông tin giảng viên thành công!")
		fmt.Printf("%+v\n", *teacher)
		utils.ReadInput("Nhấn Enter để tiếp tục...")
	} else {
		fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
		utils.ReadInput("Nhấn Enter để tiếp tục")
	}
}

func showTeacherList() {
	fmt.Println("-=-=-=-=-=- Danh sách giảng viên -=-=-=-=-=-")
	if len(teacherList) == 0 {
		fmt.Println(" ==> Không có giảng viên nào trong danh sách.")
		utils.ReadInput("Nhấn Enter để tiếp tục")
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

// Search Teacher by ID
// func searchTeacherByID() {
// 	id := utils.GetPositiveIntInput("Nhập mã số giảng viên cần tìm: ")
// 	for _, teacher := range teacherList {
// 		if teacher.ID == id {
// 			fmt.Println("Thông tin giảng viên:")
// 			fmt.Println(getInfo(teacher))
// 			return
// 		} else {
// 			fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
// 		}
// 	}
// 	return
// }

func TeacherMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("\n-=-=-=-=-=- Quản lý giảng viên -=-=-=-=-=-")
		fmt.Println("1️⃣. Thêm giảng viên")
		fmt.Println("2️⃣. Xóa giảng viên")
		fmt.Println("3️⃣. Sửa giảng viên")
		fmt.Println("4️⃣. Hiển thị danh sách giảng viên")
		fmt.Println("5️⃣. Tìm kiếm giảng viên theo mã số")
		fmt.Println("6️⃣. Quay lại menu chính 🔙")
		choice := utils.GetPositiveIntInput("Chọn một tùy chọn (1-5): ")
		switch choice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			updateTeacher()
		case 4:
			showTeacherList()
		case 5:
			// searchTeacherByID()
			id := utils.GetPositiveIntInput("Nhập mã số giảng viên cần tìm: ")
			teacher := searchTeacherByID(id)
			if teacher != nil {
				fmt.Println("Thông tin giảng viên:")
				fmt.Println(getInfo(*teacher))
				utils.ReadInput("Nhấn Enter để tiếp tục")
			} else {
				fmt.Println("Không tìm thấy giảng viên với mã số đã cho.")
				utils.ReadInput("Nhấn Enter để tiếp tục")
			}
		case 6:
			return
		default:
			fmt.Println("Tùy chọn không hợp lệ. Vui lòng chọn lại.")
		}
	}
}
