// Demo of filtering
package main

import (
    "fmt"
    "os"
    "time"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

var packetCount = 0
var snapshotLen = int32(65000)

func main() {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        fmt.Println("Error finding devices:  ", err)
        os.Exit(1)
    }
    devName := devices[0].Name
    handle, err := pcap.OpenLive(devName, 65000, false, -1 * time.Second)
    if err != nil {
        fmt.Println("Error opening device:  ", err)
        os.Exit(1)
    }
    defer handle.Close()

    // Set filter using BPF - Berkeley Packet Filter
    // Note, there are lots of online resources showing BPF syntax
    var filter string = "tcp and port 80"
    // Add filter to handle
    err = handle.SetBPFFilter(filter)
    if err != nil {
        fmt.Println("Error applying filter:  ", err)
        os.Exit(1)
    }

    fmt.Printf("Capturing %s\n", filter)

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
        packetCount++
        if packetCount > 100 {
            os.Exit(1)
        }

    }

}

