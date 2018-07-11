package main

import "fmt"

const (
	instanceCount = 20
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

	// Create intermediate subnets
	subnets := []subnet{first}
	subnets = append(subnets, generateSubnets(instanceCount)...)
	subnets = append(subnets, last)

	printLayout(subnets)

	// Fill in template with subnet info

}

func printLayout(nets []subnet) {
	fmt.Printf("%-15s %-6s %-4s %-15s\n", "left", "port", "iOS", "right")
	for i, n := range nets {
		// don't print number count on first run through
		if i > 0 {
			fmt.Printf(" %-6d %-4d ", inPort+i, i)
		}

		if n.right != "" {
			fmt.Printf("%-15s\n", n.right)
		}
		if n.left != "" {
			fmt.Printf("%-15s", n.left)
		}
	}
}
