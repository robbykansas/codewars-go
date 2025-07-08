package int32toipv4

import (
	"fmt"
)

func Int32ToIp_s(n uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", (n>>24)%256, (n>>16)%256, (n>>8)%256, n%256)
}
