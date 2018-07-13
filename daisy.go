package main

import "fmt"

const (
	instanceCount      = 4
	inPort             = 10000
	destinationAddress = "10.100.0.20"
	destinationPort    = 8080
)

type instance struct {
	count          int
	name           string
	leftNet        subnet
	rightNet       subnet
	port           int
	nextHopAddress string
	nextHopPort    int
}

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

	// Create intermediate subnets
	subnets := []subnet{first}
	subnets = append(subnets, generateSubnets(instanceCount)...)
	subnets = append(subnets, last)

	instances := createInstances(subnets)
	printInstances(instances)
	generateNacls(instances)

	// Fill in template with subnet info

}

func generateNacls(instances []instance) {
	for _, in := range instances {
		d := daisyTemplate{
			GwLeftNet:      in.leftNet.net,
			GwLeftNetmask:  in.leftNet.netmask,
			LeftAddress:    in.leftNet.left,
			GwRightNet:     in.rightNet.net,
			GwRightNetmask: in.rightNet.netmask,
			RightAddress:   in.rightNet.right,
			LeftPort:       in.port,
			NextHopAddress: in.nextHopAddress,
			NextHopPort:    in.nextHopPort,
		}
		parse(d)
		fmt.Println("################################################")
	}
}

func createInstances(nets []subnet) []instance {
	var instances []instance
	for i := 0; i < instanceCount; i++ {
		in := instance{
			count:          i + 1,
			name:           "name",
			leftNet:        nets[i],
			rightNet:       nets[i+1],
			port:           inPort + i + 1,
			nextHopAddress: nets[i+1].left,
			nextHopPort:    inPort + i + 2,
		}
		if i == instanceCount-1 {
			in.nextHopAddress = destinationAddress
			in.nextHopPort = destinationPort
		}
		instances = append(instances, in)
	}
	return instances
}

func printInstances(in []instance) {
	fmt.Printf("%-15s %-6s %-4s %-15s %s\n", "left", "port", "iOS", "right", "nextHop")
	for _, x := range in {
		fmt.Printf("%-15s %-6d %-4d %-15s %s:%d\n", x.leftNet.left, x.port, x.count, x.rightNet.right, x.nextHopAddress, x.nextHopPort)
	}
}
