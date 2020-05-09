// A simple hello world for pcap

package main

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func main() {
	version := pcap.Version()
	fmt.Println("You are using :", version)
}
