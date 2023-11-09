package main

import (
	"fmt"
	"github.com/billhathaway/consistentHash"
)

func main() {
	ch := consistentHash.New()
	ch.Add("server1")
	ch.Add("server2")
	ch.Add("server3")
	keys := []string{"A", "B", "C", "D", "E", "F", "G"}
	fmt.Println("3 servers")
	for _, key := range keys {
		server, _ := ch.Get([]byte(key))
		fmt.Printf("key=%s server=%s\n", key, server)
	}
	fmt.Println("Removing server3")
	ch.Remove("server3")
	for _, key := range keys {
		server, _ := ch.Get([]byte(key))
		fmt.Printf("key=%s server=%s\n", key, server)
	}

	fmt.Println("Removing server1")
	ch.Remove("server1")
	for _, key := range keys {
		server, _ := ch.Get([]byte(key))
		fmt.Printf("key=%s server=%s\n", key, server)
	}
}
