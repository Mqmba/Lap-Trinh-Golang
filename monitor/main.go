// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"context"
	"sync"
	"time"

	"mamba.com/monitor/processor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go processor.RunMonitor(ctx, &wg)

	time.Sleep(60 * time.Second)
	cancel()
	wg.Wait()

}
