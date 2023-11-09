package main

import (
	"fmt"
	"github.com/golang/groupcache/consistenthash"
)

func main() {
	// Cria um novo objeto consistenthash com 50 réplicas virtuais por nó
	hash := consistenthash.New(50, nil)

	// Adiciona alguns nós ao hash
	hash.Add("server1", "server2", "server3")

	// Obtém o nó para a chave especificada
	key := "minha_chave"
	server := hash.Get(key)
	fmt.Printf("A chave '%s' é mapeada para o servidor '%s'\n", key, server)
}
