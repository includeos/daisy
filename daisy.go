package main

import "fmt"

const (
	instanceCount = 3
	inPort        = 10000
)

func main() {
	// TODO: Read config file

	// TODO: Process first and last file
	first := subnet{
		net:     "10.100.0.0",
		netmask: "255.255.255.128",
		left:    "10.100.0.40",
	}
	last := subnet{
		net:     "10.100.0.0",
		netmask: "255.255.255.128",
		right:   "10.100.0.41",
	}
	// TODO: Process last destination and port

	// Create intermediate subnets
	subnets := []subnet{first}
	subnets = append(subnets, generateSubnets(instanceCount)...)
	subnets = append(subnets, last)

	printLayout(subnets)
	generateNacls(subnets)

	// Fill in template with subnet info

}

func printLayout(nets []subnet) {
	fmt.Printf("%-15s %-6s %-4s %-15s\n", "left", "port", "iOS", "right")
	for i := 0; i < instanceCount; i++ {
		fmt.Printf("%-15s %-6d %-4d %-15s\n", nets[i].left, inPort+i+1, i+1, nets[i+1].right)
	}
}

func generateNacls(nets []subnet) {
	for i := 0; i < instanceCount; i++ {
		d := daisyTemplate{
			GwLeftNet:      nets[i].net,
			GwLeftNetmask:  nets[i].netmask,
			LeftAddress:    nets[i].left,
			GwRightNet:     nets[i+1].net,
			GwRightNetmask: nets[i+1].netmask,
			RightAddress:   nets[i+1].right,
			LeftPort:       inPort + i + 1,
			NextHopAddress: nets[i+1].left,
			NextHopPort:    inPort + i + 2,
		}
		parse(d)
		fmt.Println("################################################")
	}
}
