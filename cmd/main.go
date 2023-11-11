package main

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"go-shard/internal/application/core/domain"
)

func main() {
	keyDistribution := make(map[string]int)

	shards := make([]string, 0)
	for i := 0; i < 10; i++ {
		shards = append(shards, fmt.Sprintf("Shard%d", i))
	}

	rz := domain.New(shards, xxhash.Sum64String)

	for i := 0; i < 300000; i++ {
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
	fmt.Println()

	rz.Remove("Shard1")
	keyDistribution = make(map[string]int)

	for i := 0; i < 300000; i++ {
		keyName := fmt.Sprintf("Key%d", i)
		shard := rz.Lookup(keyName)
		keyDistribution[shard]++
	}

	// -------------------------------------------------------------------------

	for shard, count := range keyDistribution {
		fmt.Printf("%s has %d keys\n", shard, count)
	}
}
