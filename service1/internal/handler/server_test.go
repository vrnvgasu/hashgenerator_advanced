package handler

import (
	hashcomposeMocks "service1/mocks/composer/hashcompose"
	"service1/pkg/composer/hashcompose"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var expectedSha256Hash = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

func TestGen(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockServer := hashcomposeMocks.NewMockHashComposeService_GenServer(mockCtrl)
	mockServer.EXPECT().Send(&hashcompose.HashObj{
		Hash: expectedSha256Hash,
	})

	list := hashcompose.StringList{List: []string{"test"}}

	server := HashComposeServiceServer{}
	err := server.Gen(&list, mockServer)
	assert.Nil(t, err)
}

func TestHash(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	mockCtrl := gomock.NewController(t)
	mockServer := hashcomposeMocks.NewMockHashComposeService_GenServer(mockCtrl)
	mockServer.EXPECT().Send(&hashcompose.HashObj{
		Hash: expectedSha256Hash,
	})

	hash("test", mockServer, &wg)
}
