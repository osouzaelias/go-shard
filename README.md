# go-shard

## Compilando arquivos .proto

Instale os pacotes para compilar arquivos .proto

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Dentro do diretório `pkg/grpc` execute o comando abaixo:

```shell
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative ./shard.proto 
```


## Referencias

- [Consistent Hashing: Algorithmic Tradeoffs](https://dgryski.medium.com/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8)
- [Client lib for redis](https://github.com/redis/go-redis/blob/21bd40a47e56e61c0598ea1bdf8e02e67d1aa651/ring.go#L28) 
- [Rendezvous hashing](https://en.wikipedia.org/wiki/Rendezvous_hashing)
- [xxHash non-cryptographic hash algorithm](https://xxhash.com/)
- [Xorshift random number generators,](https://en.wikipedia.org/wiki/Xorshift)
