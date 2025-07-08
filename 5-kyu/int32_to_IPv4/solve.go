package int32toipv4

import (
	"fmt"
	"net"
)

func Int32ToIp(n uint32) string {
	ip := make(net.IP, 4)
	ip[0] = byte(n >> 24)
	ip[1] = byte(n >> 16)
	ip[2] = byte(n >> 8)
	ip[3] = byte(n)

	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
