package main

import (
	"fmt"
	"slices"
)

func main() {

	//Khởi tạo slice rỗng
	// var numbers []int
	// fmt.Println(numbers)

	// slice := []int{10, 20, 30, 40, 50} //Kích thước tự động
	// fmt.Println(slice)

	// array := [5]int{1, 2, 3, 4, 5} //Kích thước cố định
	// fmt.Println(array)

	// fmt.Println("Slice có phải là Slice không? ", reflect.TypeOf(slice).Kind() == reflect.Slice)
	// fmt.Println("Array có phải là Array không? ", reflect.TypeOf(array).Kind() == reflect.Array)

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:4] //Tạo slice từ mảng
	// fmt.Println("Mảng ban đầu:", arr)
	// fmt.Println("arr có phải là array không?", reflect.TypeOf(arr).Kind() == reflect.Array)
	// fmt.Println("Slice được tạo từ mảng:", slice)
	// fmt.Println("slice có phải là slice không?", reflect.TypeOf(slice).Kind() == reflect.Slice)

	// slice := make([]int, 2, 5) //Tạo slice với hàm make
	// fmt.Println("Slice:", slice)
	// slice = append(slice, 60, 70, 80, 90) //Thêm phần tử vào slice
	// fmt.Println("Sau khi thêm phần tử:")
	// fmt.Println("Độ dài của slice:", len(slice))
	// fmt.Println("Sức chứa của slice:", cap(slice))
	// fmt.Println("Slice:", slice)

	// apple := []string{"Red", "Green", "Yellow"}
	// for i, color := range apple {
	// 	fmt.Printf("Apple %d: %s\n", i+1, color)
	// }

	// student := [][]string{
	// 	{"Alice", "A"},
	// 	{"Bob", "B"},
	// 	{"Charlie", "C"},
	// }
	// for _, s := range student {
	// 	fmt.Printf("Name: %s, Grade: %s\n", s[0], s[1])
	// }

	// slice := []int{10, 20, 30, 40, 50}
	// fmt.Println("Slice ban đầu:", slice)
	// fmt.Println("slice có phải là slice không?", reflect.TypeOf(slice).Kind() == reflect.Slice)
	// fmt.Println("Độ dài của slice:", len(slice))
	// fmt.Println("Sức chứa của slice:", cap(slice))
	// fmt.Println()

	// subSlice := slice[1:4] //Tạo subslice từ slice
	// fmt.Println("SubSlice:", subSlice)
	// fmt.Println("Subslice có phải là slice không?", reflect.TypeOf(subSlice).Kind() == reflect.Slice)
	// fmt.Println("Độ dài của subslice:", len(subSlice))
	// fmt.Println("Sức chứa của subslice:", cap(subSlice))

	// Tìm hiểu về package slices
	// Cloning a slice
	copied := slices.Clone([]int{1, 2, 3, 4, 5})
	fmt.Println("Slice ban đầu:", copied)

	// So sánh 2 slice
	equal := slices.Equal([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println("2 slice có bằng nhau không?", equal)

	// Tìm phần tử trong slice
	index := slices.Index([]int{10, 20, 30, 40, 50}, 30)
	fmt.Println("Vị trí phần tử 30 trong slice:", index)

	// Kiểm tra phần tử có trong slice không
	contains := slices.Contains([]int{10, 20, 30, 40, 50}, 20)
	fmt.Println("Slice có chứa phần tử 20 không?", contains)

	// Chèn thêm phần tử vào slice
	newSlice := slices.Insert([]int{1, 2, 4, 5}, 2, 3)
	fmt.Println("Slice sau khi chèn phần tử 3 vào vị trí 2:", newSlice)

	// Xoá phần tử khỏi slice
	removed := slices.Delete([]int{1, 2, 3, 4, 5}, 2, 3)
	fmt.Println("Slice sau khi xoá phần tử từ vị trí 2 đến 3:", removed)

	// Sắp xếp slice
	unsorted := []int{5, 2, 4, 1, 3}
	slices.Sort(unsorted)
	fmt.Println("Slice sau khi sắp xếp:", unsorted)

	// Đảo ngược slice
	reversed := []int{1, 2, 3, 4, 5}
	slices.Reverse(reversed)
	fmt.Println("Slice sau khi đảo ngược:", reversed)

	// Sắp xếp slice theo điều kiện tuỳ chỉnh
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slices.SortFunc(s, func(a, b int) int {
		return b - a // Sắp xếp số chẵn trước số lẻ
	})
	fmt.Println("Slice sau khi sắp xếp theo điều kiện tuỳ chỉnh:", s)

	// Min và Max trong slice
	nums := []int{42, 17, 68, 5, 99, 23}
	min := slices.Min(nums)
	max := slices.Max(nums)
	fmt.Println("Slice:", nums)
	fmt.Println("Giá trị nhỏ nhất trong slice:", min)
	fmt.Println("Giá trị lớn nhất trong slice:", max)
}
