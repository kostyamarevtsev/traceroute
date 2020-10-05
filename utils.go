package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func GetICMPMsg() []byte {

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("icmp message"),
		},
	}

	bmsg, err := msg.Marshal(nil)

	if err != nil {
		log.Fatal(err)
	}

	return bmsg
}

func Output(i *int, peer *net.Addr, finish *time.Duration) {
	fmt.Printf("|%-3d|%-16v|%-12s|\n", *i, *peer, *finish)
}
