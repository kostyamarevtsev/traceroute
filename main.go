package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
)

var target *net.IPAddr

func init() {
	dst, err := net.ResolveIPAddr("ip4", os.Args[1])
	if err != nil {
		fmt.Println("Unknown host")
	}
	target = dst
}

func main() {

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("listen err, %s", err)
	}
	defer conn.Close()

	icmpReq := GetICMPMsg()
	buf := make([]byte, 1500)

	for i := 1; ; i++ {
		start := time.Now()
		conn.IPv4PacketConn().SetTTL(i)
		err := conn.SetDeadline(time.Now().Add(time.Second * 10))
		if err != nil {
			log.Fatalf("SetDeadline err, %s", err)
		}

		if _, err := conn.WriteTo(icmpReq, target); err != nil {
			log.Fatalf("WriteTo err, %s", err)
		}

		readed, peer, err := conn.ReadFrom(buf)

		if err != nil {
			log.Fatalf("ReadFrom err, %s", err)
		}

		icmp.ParseMessage(1, buf[:readed])

		finish := time.Since(start)

		Output(&i, &peer, &finish)

		if net.ParseIP(peer.String()).Equal(target.IP) {
			break
		}

	}
}
