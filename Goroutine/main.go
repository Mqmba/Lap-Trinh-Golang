package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc \n", id)
}

func main() {
	// // defer chạy cuối cùng
	// defer fmt.Println("Kết thúc")
	// fmt.Println("Bắt đầu")
	// fmt.Println("Running 1")
	// fmt.Println("Running 2")
	// fmt.Println("Running 3")

	// start := time.Now()
	// // 4 task chạy cùng lúc khi sử dụng goroutine
	// go task(1)
	// go task(2)
	// go task(3)
	// go task(4)

	// time.Sleep(2 * time.Second)
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

	fmt.Println("Tong thời gian hoàn thành: ", time.Since(start))
}
