package main

import (
	"fmt"
	"log"

	"foreverstore/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("doing some logic with the peer outside of TCPTransport")
	// peer.Close()
	return nil
}

func main() {
	fmt.Println("Hiiii!!<3")

	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshkeFunk,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	select {}
}
