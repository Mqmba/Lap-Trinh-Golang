package teacher

import "fmt"

type Teacher struct {
	ID       int
	Name     string
	subjects string
	salary   float64
	bonus    float64
}

func getInfo(teacher Teacher) string {
	return fmt.Sprintf("Id: %d || Name: %s || Subjects: %s || Salary: %.2f || Bonus: %.2f",
		teacher.ID, teacher.Name, teacher.subjects, teacher.salary, teacher.bonus)
}

func getID(teacher Teacher) int {
	return teacher.ID
}

func checkDuplicateID(id int, list []Teacher) bool {
	for _, teacher := range list {
		if getID(teacher) == id {
			return true
		}
	}
	return false
}
