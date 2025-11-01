package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct
type Person struct {
	Name   string `json:"Tên"`
	Gender int    `json:"Giới tính"`
	Phone  int    `json:"Số điện thoại"`
}

func (p Person) showInfo() {
	fmt.Println("Name: ", p.Name)
	fmt.Println("Gender: ", p.Gender)
	fmt.Println("Phone: ", p.Phone)
}

func (p Person) clear() {
	p.Name = ""
	p.Gender = 0
	p.Phone = 0
}

func main() {
	huy := Person{
		Name:   "Huy",
		Gender: 1,
		Phone:  123456,
	}
	huy.showInfo()
	fmt.Println("-----")
	huy.clear()
	huy.showInfo()
	fmt.Println("-----")

	// Tony := Person{
	// 	Name:   "Tony",
	// 	Gender: 1,
	// 	Phone:  654321,
	// }
	// Tony.showInfo()
	// fmt.Println("-----")

	output, err := json.Marshal(huy)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("JSON Output:", string(output))
	fmt.Println("-----")
}
