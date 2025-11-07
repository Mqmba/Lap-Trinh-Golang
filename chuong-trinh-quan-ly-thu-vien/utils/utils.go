package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetIntInput(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err == nil && value >= 0 {
			return value
		}
		fmt.Println("Vui lòng nhập một số lớn hơn 0.")
	}
}

func GetPositiveIntInput(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui lòng nhập một số nguyên dương.")
	}
}

func GetPositiveFloatInput(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value >= 0 {
			return value
		}
		fmt.Println("Vui lòng nhập một số thực dương.")
	}
}

func GetOptionalPositiveFloatInput(prompt string, oldValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	value, err := strconv.ParseFloat(input, 64)
	if err == nil && value >= 0 {
		return value
	}
	fmt.Println("Giá trị không hợp lệ. Giữ nguyên giá trị cũ.")
	return oldValue
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("Vui lòng nhập một chuỗi không rỗng.")
	}
}

func GetOptionalString(prompt string, oldValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}
	return input
}

func ClearScreen() {
	// fmt.Print("\033[H\033[2J")
	var cmd *exec.Cmd
	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Lỗi khi xóa màn hình:", err)
	}
}

type HasId interface {
	GetID() int
}

func CheckDuplicateID[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetID() == id {
			return true
		}
	}
	return false
}
