package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"log"
)

func main() {
	for {
		fmt.Println("input shards number:")
		var numShards int64
		if _, err := fmt.Scanf("%d", &numShards); err != nil {
			log.Fatal(err)
		}

		fmt.Println("input key:")
		var inputKey string
		if _, err := fmt.Scanf("%s", &inputKey); err != nil {
			log.Fatal(err)
		}

		// Chave a ser sharded
		key := []byte(inputKey)

		// Calcular o hash da chave usando MurmurHash3
		hashValue := murmur3.Sum32(key)

		// Determinar o shard para a chave
		shard := hashValue % uint32(numShards)

		fmt.Printf("A chave '%s' pertence ao shard número %d.\n", key, shard)
	}
}
