// Demo showing how to open a device
package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "os"
    "time"
)

func main() {


    devices, err := pcap.FindAllDevs()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }


    devName := devices[0].Name
    // OpenLive returns a handle
    handle, err := pcap.OpenLive(
            devName,
            65000,
            false,
            30 * time.Second,
            )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer handle.Close()

    // Create new packet source from handle
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }


}

