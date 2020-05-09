// How to decode packet layers
package main

import (
    "fmt"
    "os"
    "time"
    "log"
    "strings"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/layers"
)

var (
    snapLen int32 = 1024
    promisc bool = false
    timeout time.Duration = -1 * time.Second
    devName string
    filter string
    handle *pcap.Handle
    err error
    packetCount int
    packetSource *gopacket.PacketSource
)


func main() {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }

    devName = devices[0].Name
    handle, err = pcap.OpenLive(devName, snapLen, promisc, timeout)
    if err != nil {
        log.Fatal(err)
    }
    
    // Set a filter
    filter = "tcp and port 80 or port 443"
    err = handle.SetBPFFilter(filter)
    if err != nil {
        log.Fatal(err)
    }

    defer handle.Close()
    
    fmt.Printf("Capturing %s\n", filter)
    // Create new packetsource
    packetSource = gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {

        fmt.Println("***********************************")
        printPacketInfo(packet) 
        packetCount++
        
        if packetCount > 10 {
            os.Exit(1)
        }


    }
}

func printPacketInfo(packet gopacket.Packet) {
    ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
    if ethernetLayer != nil {
        fmt.Println("Ethernet layer detected!")
        ethernetPacket := ethernetLayer.(*layers.Ethernet)
        fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
        fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
        // Ethernet is typically IPv$ but could be ARP or other
        fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
        fmt.Println()
    }

    // Check for IP layer (even though ethernet tells us)
    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        fmt.Println("IPv4 layer detected!")
        ip := ipLayer.(*layers.IPv4)
        fmt.Println("Source IP: ", ip.SrcIP)
        fmt.Println("Destination IP: ", ip.DstIP)
        fmt.Println("Protocol: ", ip.Protocol)
        fmt.Println()
    }
    
    // Test if packet is TCP
    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer != nil {
        fmt.Println("TCP Transport layer detected!")
        tcp := tcpLayer.(*layers.TCP)
        fmt.Println("Source port: ", tcp.SrcPort)
        fmt.Println("Destination port: ", tcp.DstPort)
        fmt.Println("Sequence N: ", tcp.Seq)
        fmt.Println()
    }
    // Iterate over the packets
    fmt.Println("All packet layers")
    for _, layer := range packet.Layers() {
        fmt.Println("- ", layer.LayerType())
    }

    // Any payload layers listed in the loop immediately above
    // is the same as the application layer
    appLayer := packet.ApplicationLayer()
    if appLayer != nil {
        fmt.Println("Application layer / payload found")
        fmt.Println(string(appLayer.Payload()))

        // Search for a string inside the payload
        if strings.Contains(string(appLayer.Payload()), "HTTP") {
            fmt.Println("HTTP found!")
        }
    }

    // Check for errors
    if err := packet.ErrorLayer(); err != nil {
        fmt.Println("Error decoding some part of the packet:", err)
    }


}


