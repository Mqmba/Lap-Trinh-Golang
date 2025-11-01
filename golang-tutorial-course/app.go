package main

// // Biến toàn cục
// var address string = "123 Main St"

// var (
// 	city    = "Hanoi"
// 	country = "Vietnam"
// )

func main() {

	// // fmt.Println("Hello, Mamba!")
	// // randomUser()

	// // Khai bao bien su dung từ khóa var
	// var fullName = "Nguyễn Hữu Mamba"
	// fmt.Println("Full Name:", fullName)

	// var age = 30
	// fmt.Println("Age:", age)

	// // Khai bao biến không sử dụng var
	// phone := "0123456789"
	// fmt.Println("Phone:", phone)

	// fmt.Println("Address:", address)

	// // Khai báo hằng số
	// const isMarried = false
	// if isMarried {
	// 	fmt.Println("Married")
	// } else {
	// 	fmt.Println("Not Married")
	// }
	// // Bài tập tính điểm trung bình
	// toan, ly, hoa := 9, 8, 7
	// diemTrungBinh := float64(toan+ly+hoa) / 3
	// fmt.Println("Điểm trung bình:", diemTrungBinh)

	// fmt.Println("City:", city)
	// fmt.Println("Country:", country)

	// var monan, gia string
	// fmt.Println("Nhập món ăn yêu thích của bạn:")
	// fmt.Scan(&monan, &gia)
	// fmt.Printf("Món ăn yêu thích của bạn là: %s, giá: %s\n", monan, gia)

	// var name string
	// var age int
	// fmt.Println("Nhập tên của bạn:")
	// fmt.Scanf("%[^\n]", &name)
	// fmt.Println("Nhập tuổi của bạn:")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("Tên của bạn là: %s, tuổi: %d \n", name, age)

	// message := fmt.Sprint("Hello ", "Mamba!")
	// fmt.Println(message)

	// ten := "Nguyen Huu Mamba"
	// tuoi := 30
	// chieucao := 1.75
	// daTotNghiep := true
	// phanTram := 99.5
	// ketqua := fmt.Sprintf("Ten: %s, Tuoi: %d, Chieu cao: %.2f, Da tot nghiep: %t, Phan tram: %.1f%%", ten, tuoi, chieucao, daTotNghiep, phanTram)
	// fmt.Println(ketqua)

	// s1 := 15
	// s2 := 4

	// tong := s1 + s2
	// hieu := s1 - s2
	// tich := s1 * s2
	// thuong := s1 / s2
	// du := s1 % s2

	// fmt.Printf("Tổng: %d + %d = %d\n", s1, s2, tong)
	// fmt.Printf("Hiệu: %d - %d = %d\n", s1, s2, hieu)
	// fmt.Printf("Tích: %d * %d = %d\n", s1, s2, tich)
	// fmt.Printf("Thương: %d / %d = %d\n", s1, s2, thuong)
	// fmt.Printf("Dư: %d %% %d = %d\n", s1, s2, du)

	// var n, m int
	// fmt.Println("Nhập vào hai số nguyên n và m:")
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Lỗi khi đọc input:", err)
	// 	return
	// }
	// if n < m {
	// 	fmt.Println("n < m")
	// } else if n > m {
	// 	fmt.Println("n > m")
	// } else {
	// 	fmt.Println("n = m")
	// }

	// var n int
	// fmt.Scanf("%d", &n)
	// switch n {
	// case 1:
	// 	fmt.Println("Thứ Hai")
	// case 2:
	// 	fmt.Println("Thứ Ba")
	// case 3:
	// 	fmt.Println("Thứ Tư")
	// case 4:
	// 	fmt.Println("Thứ Năm")
	// case 5:
	// 	fmt.Println("Thứ Sáu")
	// case 6:
	// 	fmt.Println("Thứ Bảy")
	// case 7:
	// 	fmt.Println("Chủ Nhật")
	// default:
	// 	fmt.Println("Ngày không hợp lệ")
	// }

	// Bài tập in các số từ 1 đến 100, trừ các số 6, 48, 75, 89
	// var first bool = true
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}
	// 	if !first {
	// 		fmt.Print(", ")
	// 	}
	// 	fmt.Print(i)
	// 	first = false
	// }

	// var xhtml = ""
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}
	// 	xhtml += strconv.Itoa(i)
	// 	if i != 100 {
	// 		xhtml += ", "
	// 	} else {
	// 		xhtml += "\n"
	// 	}
	// }
	// fmt.Print(xhtml)

	// Bài tập in các số lẻ từ 1 đến 100
	// Cứ 3 số in một dòng
	// Số đầu tiên không in dấu phẩy
	// Số cuối cùng không in dấu phẩy
	// count := 0
	// for j := 1; j <= 100; j++ {
	// 	if j%2 == 0 {
	// 		continue
	// 	}
	// 	if count%3 != 0 {
	// 		fmt.Print(", ")
	// 	}
	// 	fmt.Print(j)
	// 	count++
	// 	if count%3 == 0 {
	// 		fmt.Println()
	// 	}
	// }

}
