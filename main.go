package main

import (
	"go-distributed-storage/server"
	"go-distributed-storage/storage"
)

func main() {
	s := server.New(storage.New())
	s.Start()
}
