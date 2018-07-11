package main

import (
	"net"

	"github.com/brotherpowers/ipsubnet"
)

type subnet struct {
	net     string
	netmask string
	first   string
	last    string
	left    string
	right   string
}

var (
	network = net.ParseIP("10.90.0.0")
	netmask = 16
)

func generateSubnets(num int) []subnet {
	subs := []subnet{}
	IP := network
	for i := 0; i < num-2; i++ {
		s := generateSubnet(IP.String())
		subs = append(subs, s)
		IP = nextIP(IP, 4)
	}
	// finally append the last subnet
	subs = append(subs, lastSubnet())
	return subs
}

// lastSubnet finds the last subnet in the block
func lastSubnet() subnet {
	sub := ipsubnet.SubnetCalculator(network.String(), netmask)
	lastAddress := sub.GetIPAddressRange()[1]
	subLast := ipsubnet.SubnetCalculator(lastAddress, 30)
	s := generateSubnet(subLast.GetNetworkPortion())

	return s
}

func generateSubnet(start string) subnet {
	sub := ipsubnet.SubnetCalculator(start, 30)
	s := subnet{}
	s.net = start
	s.netmask = sub.GetSubnetMask()
	s.first = nextIP(net.ParseIP(start), 1).String()
	s.last = nextIP(net.ParseIP(s.first), 1).String()

	s.right = s.first
	s.left = s.last
	return s
}

func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}
