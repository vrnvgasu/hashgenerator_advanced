package hashhandler

import (
	"context"
	"service2/pkg/composer/hashcompose"
	"time"

	"google.golang.org/grpc"
)

func Generate(list []string) ([]string, error) {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	//conn, err := grpc.DialContext(cwt, "pdf-compose-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	uc := hashcompose.NewHashComposeServiceClient(conn)

	stream, err := uc.Gen(cwt, &hashcompose.StringList{List: list})
	if err != nil {
		panic(err)
	}

	grpcHashes := make([]string, 0, len(list))
	for {
		grpcRes, err := stream.Recv() // вынимает порцию данных из стрима
		if grpcRes == nil {
			break
		}
		if err != nil {
			panic(err)
		}

		grpcHash := grpcRes.GetHash()
		grpcHashes = append(grpcHashes, grpcHash)
	}

	return grpcHashes, nil
}
