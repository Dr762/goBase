package network

import (
	"net"
	"fmt"
	"log"
)

func GetMask(ipAddr string) {

	addr := net.ParseIP(ipAddr)
	if addr == nil {
		fmt.Println("Invalid address")
	}

	mask := addr.DefaultMask()
	nt := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address:", addr)
	fmt.Println("Mask(hex):", mask.String())
	fmt.Println("Length:", bits)
	fmt.Println("Leading ones count:", ones)
	fmt.Println("Network:", nt.String())

}

func ResolveIP(hostname string) {
	addr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resolved address is", addr)
}

func HostLookup(hostname string) {
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
}

func PortLookup(netType string, service string) {
	port, err := net.LookupPort(netType, service)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service port", port)
}

func Ping(host string) {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	if err != nil {
		log.Fatal(err)
	}

	var msg [512] byte

	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 13 //identifier[1]msg[6] = 0
	msg[6] = 0 // sequence[0]
	msg[7] = 37 // sequence[1]
	length := 8

	check := checkSum(msg[0:length])
	msg[2] = byte(check >> 0)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:length])
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Read(msg[0:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	//even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n]+1)
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16

	var answer uint16 = uint16(^sum)

	return answer
}
