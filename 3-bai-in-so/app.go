package main

import (
	"fmt"
	"strconv"
)

func main() {
	// // 1から100までの整数のうち、3の倍数を除いた数値をカンマ区切りで表示する
	// // 例: 1, 2, 4, 5, 7, 8, ..., 98, 100
	// // ヒント: continue文を使うと良い
	// // ポイント: 最後の数値の後にカンマを表示しないようにする
	// var point = true
	// var str string = ""
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		continue
	// 	}
	// 	if !point {
	// 		str += ", "
	// 	}
	// 	// Convert int to string
	// 	// https://pkg.go.dev/strconv#Itoa
	// 	// Itoa は整数を10進数の文字列に変換します。
	// 	str += strconv.Itoa(i)
	// 	point = false
	// }
	// fmt.Println(str)

	// // 1から100までの整数のうち、偶数を除いた数値を3つずつ改行して表示する
	// // 例:
	// // 1, 3, 5,
	// count := 0
	// str := ""
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	if count%3 == 0 {
	// 		str += "\n"
	// 	} else {
	// 		str += ", "
	// 	}
	// 	str += strconv.Itoa(i)
	// 	count++
	// }
	// fmt.Print(str)

	// 2つの整数を入力し、その範囲の九九を表示する
	// 例:
	// 整数1: 3
	var num1, num2 int
	str := ""
	fmt.Println("2つの整数を入力してください")
	fmt.Print("整数1: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Print("整数2: ")
	fmt.Scanf("%d\n", &num2)
	if num1 == 0 || num2 == 0 {
		fmt.Println("0は入力しないでください")
		return
	} else if num1 < 0 || num2 < 0 {
		fmt.Println("負の数は入力しないでください")
		return
	} else if num1 > num2 {
		fmt.Println("整数1は整数2以下の数を入力してください")
		return
	} else {
		for i := num1; i <= num2; i++ {
			str += strconv.Itoa(i) + "の段の九九:\n"
			for j := 1; j <= 10; j++ {
				str += strconv.Itoa(i) + " x " + strconv.Itoa(j) + " = " + strconv.Itoa(i*j) + "\n"
				if j == 10 {
					str += "\n"
				}
			}
		}
		fmt.Print(str)
	}
}
