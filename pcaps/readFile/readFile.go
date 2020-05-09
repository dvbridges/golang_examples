// Demo showing how to read a pcap file
package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "os"
)

func main() {
    // Open file
    handle, err := pcap.OpenOffline("./dump.pcap")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer handle.Close()

    // Create packetsource
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }

}

