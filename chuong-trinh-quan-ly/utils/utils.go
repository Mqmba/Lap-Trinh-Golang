package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
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
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui lòng nhập một số thực dương.")
	}
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
