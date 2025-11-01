package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hinhchunhat struct {
	chieuDai  float32
	chieuRong float32
}

// func inputHinhchunhat() Hinhchunhat {
// 	var hinhchunhat Hinhchunhat
// 	for {
// 		fmt.Print("Nhap chieu dai: ")
// 		_, err := fmt.Scanf("%f", &hinhchunhat.chieuDai)
// 		if err == nil && hinhchunhat.chieuDai > 0 {
// 			break
// 		}
// 		fmt.Println("Loi khi nhap chieu dai")
// 	}
// 	for {
// 		fmt.Print("Nhap chieu rong: ")
// 		_, err := fmt.Scanf("%f", &hinhchunhat.chieuRong)
// 		if err == nil && hinhchunhat.chieuRong > 0 {
// 			break
// 		}
// 		fmt.Println("Loi khi nhap chieu rong")
// 	}
// 	return hinhchunhat
// }

// Cách 2: Sử dụng bufio và strconv để đọc và chuyển đổi dữ liệu
// Nhập dữ liệu từ bàn phím
func inputHinhchunhatV2() Hinhchunhat {
	var cd, cr float32
	for {
		var err error
		cd, err = readFloat("Nhap chieu dai: ")
		if err != nil || cd <= 0 {
			fmt.Println("Loi khi nhap chieu dai")
		} else {
			break
		}
	}

	for {
		var err error
		cr, err = readFloat("Nhap chieu rong: ")
		if err != nil || cr <= 0 {
			fmt.Println("Loi khi nhap chieu rong")
		} else {
			break
		}
	}

	return Hinhchunhat{
		chieuDai:  cd,
		chieuRong: cr,
	}
}

func readFloat(prompt string) (float32, error) {
	// Prompt user for input
	fmt.Print(prompt)

	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("Error reading input: %v", err)
	}

	// Trim whitespace and parse float
	input = strings.TrimSpace(input)

	// Convert string to float32
	value, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, fmt.Errorf("Invalid number: %v", err)
	}

	return float32(value), nil
}

func (h *Hinhchunhat) chuVi() float32 {
	return 2 * (h.chieuDai + h.chieuRong)
}

func (h *Hinhchunhat) dienTich() float32 {
	return h.chieuDai * h.chieuRong
}

func main() {
	hinhchunhat := inputHinhchunhatV2()
	fmt.Printf("Chu vi hinh chu nhat: %.2f\n", hinhchunhat.chuVi())
	fmt.Printf("Dien tich hinh chu nhat: %.2f\n", hinhchunhat.dienTich())
}
