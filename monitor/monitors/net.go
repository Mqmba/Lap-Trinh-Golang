package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetMonitor struct {
}

func (cpu *NetMonitor) Name() string {
	return "Net"
}

func (m *NetMonitor) Check(ctx context.Context) string {
	netStat, err := net.IOCountersWithContext(ctx, false)
	if err != nil && len(netStat) == 0 {
		return "N/A"
	}

	// fmt.Printf("%+v", netStat)
	value := fmt.Sprintf("Send : %d KB, Recv : %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)

	return value
}
