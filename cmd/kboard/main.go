package main

import (
	"fmt"
	"log"

	"github.com/scgolang/kboard"
)

func main() {
	kb, err := kboard.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := kb.Open(); err != nil {
		log.Fatal(err)
	}
	pkts, err := kb.Packets()
	if err != nil {
		log.Fatal(err)
	}
	for pkt := range pkts {
		fmt.Printf("%#v\n", pkt.Data)
	}
}
