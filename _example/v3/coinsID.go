package main

import (
	"fmt"
	"log"

	gecko "github.com/nwosuudoka/go-gecko/v3"
)

// this is an example querying dogecoin
func main() {
	cg := gecko.NewClient(nil)
	coin, err := cg.CoinsID("dogecoin", true, true, true, true, true, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coin)
}
