package teacher

import "fmt"

type Teacher struct {
	ID       int
	Name     string
	Subjects string
	Salary   float64
	Bonus    float64
}

func getInfo(teacher Teacher) string {
	return fmt.Sprintf("Id: %d || Name: %s || Subjects: %s || Total Compensation: %.2f",
		teacher.ID, teacher.Name, teacher.Subjects, getTotalCompensation(teacher))
}

func getTotalCompensation(teacher Teacher) float64 {
	return teacher.Salary + teacher.Bonus
}

func (t Teacher) GetID() int {
	return t.ID
}

// func getID(teacher Teacher) int {
// 	return teacher.ID
// }
//
// func checkDuplicateID(id int, list []Teacher) bool {
// 	for _, teacher := range list {
// 		if getID(teacher) == id {
// 			return true
// 		}
// 	}
// 	return false
// }
