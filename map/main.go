package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	// Map với struct
	// employees := map[string]Employee{
	// 	"e1": {"John", 35, "Worker"},
	// 	"e2": {"Huy", 25, "Developer"},
	// 	"e3": {"Hung", 30, "IT Helpdesk"},
	// }
	// // fmt.Printf("%+v\n", employees)

	// for key, value := range employees {
	// 	fmt.Printf("Giá trị của key %s\n", key)
	// 	fmt.Printf("Name: %s\n", value.Name)
	// 	fmt.Printf("Age: %d\n", value.Age)
	// 	fmt.Printf("Hung: %s\n", value.Role)
	// 	fmt.Println()
	// }

	//Map với slice
	studentSubjects := map[string][]string{
		"Huy":  {"Java", "Go"},
		"Hung": {"JavaScript", "TypeScript"},
	}

	// Append slice
	// studentSubjects["Huy"] = append(studentSubjects["Huy"], "Python")
	// fmt.Printf("Môn học của Huy là : %s \n", studentSubjects["Huy"][0])
	// fmt.Printf("Môn học của Huy là : %s \n", studentSubjects["Huy"][1])

	for key, value := range studentSubjects {
		for _, subjects := range value {
			fmt.Printf("Môn học của %s là : %s\n", key, subjects)
		}
	}

	// // Khai báo và khởi tạo trực tiếp
	// drink := map[string]int{
	// 	"coffee": 500,
	// 	"tea":    300,
	// }
	// fmt.Printf("%+v\n", drink)

	// student := map[int]string{
	// 	10: "Tuan",
	// 	15: "Teo",
	// 	20: "Ti",
	// 	25: "John",
	// }
	// fmt.Printf("%+v\n", student)

	// // Khai báo bằng make
	// m := make(map[string]int)
	// m["a"] = 1
	// m["b"] = 2
	// m["c"] = 3
	// fmt.Printf("%+v\n", m)

	// monan := make(map[string]int)
	// monan["chao"] = 500
	// monan["com"] = 300
	// fmt.Println(monan)

	// value, exists := monan["com"]
	// if exists {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("không tồn tại trong map")
	// }

	// for key, value := range student {
	// 	fmt.Printf("Student[%d] : %s\n", key, value)
	// }
}
