package hashhandler

import (
	"context"
	"service2/internal/lg"
	"service2/pkg/composer/hashcompose"
	"time"

	"google.golang.org/grpc"
)

func Generate(ctx context.Context, list []string) ([]string, error) {
	const (
		op   = "Generate"
		pack = "hashhandler"
	)

	cwt, _ := context.WithTimeout(ctx, time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		lg.Fatal(ctx, op, pack, err)
		panic(err)
	}
	defer conn.Close()

	uc := hashcompose.NewHashComposeServiceClient(conn)

	stream, err := uc.Gen(cwt, &hashcompose.StringList{List: list})
	if err != nil {
		lg.Fatal(ctx, op, pack, err)
		panic(err)
	}

	grpcHashes := make([]string, 0, len(list))
	for {
		grpcRes, err := stream.Recv() // вынимает порцию данных из стрима
		if grpcRes == nil {
			break
		}
		if err != nil {
			lg.Fatal(ctx, op, pack, err)
			panic(err)
		}

		grpcHash := grpcRes.GetHash()
		grpcHashes = append(grpcHashes, grpcHash)
	}

	return grpcHashes, nil
}
