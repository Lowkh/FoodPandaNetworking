package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interfaces: %v\n", i.Name)
		fmt.Printf("MAC address: %v\n", i.HardwareAddr)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)

		}

		addresses, err := byName.Addrs()
		if err != nil {
			fmt.Println(err)
			return
		}

		for k, v := range addresses {
			fmt.Printf("Interface Address # %v: %v\n", k, v.String())

		}
		fmt.Println()
	}

}
