package student

import "fmt"

type Student struct {
	ID    int
	Name  string
	Class string
	Score []float64
}

func getInfo(s Student) string {
	return fmt.Sprintf("Id: %d || Name: %s || Class: %s || Average Score: %.2f", s.ID, s.Name, s.Class, getAverageScore(s))
}

func getAverageScore(s Student) float64 {
	total := 0.0
	for _, score := range s.Score {
		total += score
	}
	if len(s.Score) == 0 {
		return 0
	}
	return total / float64(len(s.Score))
}

func (s Student) GetID() int {
	return s.ID
}

// func getID(student Student) int {
// 	return student.ID
// }

// func checkDuplicateID(id int, list []Student) bool {
// 	for _, student := range list {
// 		if getID(student) == id {
// 			return true
// 		}
// 	}
// 	return false
// }
