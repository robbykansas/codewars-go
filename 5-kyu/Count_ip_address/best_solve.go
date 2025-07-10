package countipaddress

import (
	"encoding/binary"
	"net"
)

func IpsBetween_s(first, last string) int {
	firstVal := binary.BigEndian.Uint32(net.ParseIP(first)[12:16]) // last 4 bytes because net/ip supports ipv6
	lastVal := binary.BigEndian.Uint32(net.ParseIP(last)[12:16])
	return int(lastVal) - int(firstVal)
}
