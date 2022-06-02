package LinkerGrpcAgent

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/linkerProto"
)

type GrpcAgent struct {
	LinkerGrpc linkerProto.LinkerUseCaseClient
}

func MakeAgent(gRPC linkerProto.LinkerUseCaseClient) GrpcAgent {
	return GrpcAgent{LinkerGrpc: gRPC}
}

func (agent GrpcAgent) Get(hash string) (string, error) {
	url, err := agent.LinkerGrpc.Get(context.Background(), &linkerProto.HashWrapper{Hash: hash})
	if err != nil {
		return "", err
	}
	return url.Url, nil
}

func (agent GrpcAgent) Create(url string) (string, error) {
	returnedHash, err := agent.LinkerGrpc.Create(context.Background(), &linkerProto.UrlWrapper{
		Url: url,
	})

	if err != nil {
		return "", err
	}

	return returnedHash.Hash, nil
}

func (agent GrpcAgent) Count(hash string) (int64, error) {
	fmt.Println("in agent count")
	response, err := agent.LinkerGrpc.Count(context.Background(), &linkerProto.HashWrapper{Hash: hash})
	if err != nil {
		return -1, err
	}
	return response.Result, nil
}
