// Demo showing how to create and send packets
package main

import (
    "fmt"
    "log"
    "net"
    "time"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/layers"
)

var (
    snapLen int32 = 1024
    promisc bool = false
    timeout time.Duration = 30 * time.Second
    packetCount int = 0
)

func main() {
    // Get all devs
    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }
    
    // Open connection
    devName := devices[0].Name
    handle, err := pcap.OpenLive(devName, snapLen, promisc, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    // Send some raw bytes over the wire
    rawBytes := []byte{10, 20, 30}
    err = handle.WritePacketData(rawBytes)
    if err != nil {
        fmt.Println(err)
        // log.Fatal(err)
    }

    // Create a properly formed packet, just with
    // empty details. Should fill out MAC addresses,
    //IP addresses, etc.
    buffer := gopacket.NewSerializeBuffer()
    var options gopacket.SerializeOptions
    gopacket.SerializeLayers(buffer, options,
        &layers.Ethernet{},
        &layers.IPv4{},
        &layers.TCP{},
        gopacket.Payload(rawBytes),
    )

    outgoingPacket := buffer.Bytes()
    // Send the packet
    err = handle.WritePacketData(outgoingPacket)
    if err != nil {
        log.Fatal(err)
    }

    // This time fill out some information
    // This time lets fill out some information
    ipLayer := &layers.IPv4{
        SrcIP: net.IP{127, 0, 0, 1},
        DstIP: net.IP{8, 8, 8, 8},
    }

    ethernetLayer := &layers.Ethernet{
        SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
        DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
    }

    tcpLayer := &layers.TCP{
        SrcPort: layers.TCPPort(4321),
        DstPort: layers.TCPPort(80),
    }

    // Create the packet with layers
    buffer = gopacket.NewSerializeBuffer()
    gopacket.SerializeLayers(buffer, options,
        ethernetLayer,
        ipLayer,
        tcpLayer,
        gopacket.Payload(rawBytes),
    )

    outgoingPacket = buffer.Bytes()

}
