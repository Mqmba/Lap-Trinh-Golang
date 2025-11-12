package main

import (
	"context"
	"fmt"
	"time"
)

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Công việc đã bị hủy: ", ctx.Err())
			return
		default:
			priority := ctx.Value("priority")
			fmt.Println("Đang làm việc của task ưu tiên với mức độ: ", priority)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "priority", "high")

	go employee(ctx)

	time.Sleep(3 * time.Second)
}
