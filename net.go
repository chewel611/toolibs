package toolibs

import (
	"net"
)

// GetAddrs 获取本地所有地址
func GetAddrs() (addresses []string) {
	addresses = make([]string, 2)
	inters, err := net.Interfaces()
	if err == nil {
		for _, inter := range inters {
			// 网卡状态不可用
			if inter.Flags&net.FlagUp == 0 {
				continue
			}
			// 排除本地Loopback
			if inter.Flags&net.FlagLoopback != 0 {
				continue
			}
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ip, ok := addr.(*net.IPNet); ok && ip.IP.To4() != nil {
					addresses = append(addresses, ip.IP.String())
				}
			}
		}
	}
	return
}
