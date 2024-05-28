package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"service1/pkg/composer/hashcompose"
	"sync"
)

type HashComposeServiceServer struct {
	hashcompose.HashComposeServiceServer
}

func (c *HashComposeServiceServer) Gen(list *hashcompose.StringList, stream hashcompose.HashComposeService_GenServer) error {
	stringList := list.GetList()

	var wg sync.WaitGroup
	for _, s := range stringList {
		wg.Add(1)
		go hash(s, stream, &wg)
	}

	wg.Wait()
	return nil
}

func hash(s string, stream hashcompose.HashComposeService_GenServer, wg *sync.WaitGroup) {
	h := sha256.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))

	err := stream.Send(&hashcompose.HashObj{
		Hash: hash,
	})
	if err != nil {
		log.Println("hash err:", err)
	}

	wg.Done()
}
