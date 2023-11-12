package main

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"go-shard/internal/application/core/domain"
)

func main() {
	keyDistribution := make(map[string]int)

	fmt.Println()
	fmt.Println("Adds 10 shads")

	shards := make([]string, 0)
	for i := 0; i < 10; i++ {
		shards = append(shards, fmt.Sprintf("Shard%d", i))
	}

	rz := domain.New(shards, xxhash.Sum64String)

	fmt.Println("Distributes 100,000 keys")
	fmt.Println()

	for i := 0; i < 100000; i++ {
		keyName := fmt.Sprintf("Key%d", i)
		shard := rz.Lookup(keyName)
		keyDistribution[shard]++
	}

	for shard, count := range keyDistribution {
		fmt.Printf("%s has %d keys\n", shard, count)
	}

	// -------------------------------------------------------------------------

	fmt.Println()
	fmt.Println("Remove o Shard1")

	rz.Remove("Shard1")
	keyDistribution = make(map[string]int)

	fmt.Println()
	fmt.Println("Distributes 100,000 keys")

	for i := 0; i < 100000; i++ {
		keyName := fmt.Sprintf("Key%d", i)
		shard := rz.Lookup(keyName)
		keyDistribution[shard]++
	}

	// -------------------------------------------------------------------------

	for shard, count := range keyDistribution {
		fmt.Printf("%s has %d keys\n", shard, count)
	}
}
