package main

import (
	"fmt"
	"log"

	"github.com/lutefd/distributed-file-storage/p2p"
)

func main(){
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept();
	err != nil {
		log.Fatal(err)
	}
	fmt.Println("We're gucci")
	select {}
}