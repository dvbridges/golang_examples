// A demo to find your network devices

package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("%T\n", devices) // []pcap.Inferface

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Printf("\nName: %s\n", device.Name)
		fmt.Printf("Desc: %s\n", device.Description)
		fmt.Println("Device addresses: ")
		for _, addr := range device.Addresses {
			fmt.Printf("\nDevice IP: %s\n", addr.IP)
			fmt.Printf("Netmask: %s\n", addr.Netmask)
			fmt.Printf("Broadcast addr: %s\n", addr.Broadaddr)
			fmt.Printf("P2P Dest: %s\n", addr.P2P)
		}
	}
}
