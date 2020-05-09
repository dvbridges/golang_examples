// Demo showing how to write packets to a file
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var packetCounter = 0
var snapLen = uint32(65000)

func main() {

	// Create file for writing
	file, err := os.Create("godump.pcap")
	if err != nil {
		fmt.Println(err)
	}
	writer := pcapgo.NewWriter(file)
	writer.WriteFileHeader(snapLen, layers.LinkTypeEthernet)
	defer file.Close()

	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println("Error finding devices: ", err)
		os.Exit(1)
	}

	devName := devices[0].Name
	handle, err := pcap.OpenLive(devName, 650000, false, -1*time.Second)
	if err != nil {
		fmt.Println("Error opening devices: ", err)
		os.Exit(1)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCounter++

		if packetCounter >= 100 {
			break
		}
	}

}
