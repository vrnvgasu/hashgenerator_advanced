package hashhandler

import (
	"context"
	"service2/internal/lg"
	"service2/pkg/composer/hashcompose"
	"time"

	"github.com/vrnvgasu/logwrapper"
	"google.golang.org/grpc"
)

func Generate(list []string) ([]string, error) {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	//conn, err := grpc.DialContext(cwt, "pdf-compose-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		lg.Logger.Payload(logwrapper.NewPayload().Op("Generate").Package("hashhandler")).Fatal(err)
		panic(err)
	}
	defer conn.Close()

	uc := hashcompose.NewHashComposeServiceClient(conn)

	stream, err := uc.Gen(cwt, &hashcompose.StringList{List: list})
	if err != nil {
		lg.Logger.Payload(logwrapper.NewPayload().Op("Generate").Package("hashhandler")).Fatal(err)
		panic(err)
	}

	grpcHashes := make([]string, 0, len(list))
	for {
		grpcRes, err := stream.Recv() // вынимает порцию данных из стрима
		if grpcRes == nil {
			break
		}
		if err != nil {
			lg.Logger.Payload(logwrapper.NewPayload().Op("Generate").Package("hashhandler")).Fatal(err)
			panic(err)
		}

		grpcHash := grpcRes.GetHash()
		grpcHashes = append(grpcHashes, grpcHash)
	}

	return grpcHashes, nil
}
