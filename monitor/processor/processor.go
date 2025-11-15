package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"mamba.com/monitor/monitors"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			monitor := monitors.CPUMonitor{}
			cpuPercent := monitor.Check(ctx)
			fmt.Println(cpuPercent)
		}
	}
}
