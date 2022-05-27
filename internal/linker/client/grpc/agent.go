package LinkerGrpcAgent

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/linkerProto"
)

type GrpcAgent struct {
	LinkerGrpc linkerProto.LinkerUseCaseClient
}

func MakeAgent(gRPC linkerProto.LinkerUseCaseClient) GrpcAgent {
	return GrpcAgent{LinkerGrpc: gRPC}
}

func (agent GrpcAgent) Get(url string) (string, error) {
	hash, err := agent.LinkerGrpc.Get(context.Background(), &linkerProto.UrlWrapper{Url: url})
	if err != nil {
		return "", err
	}
	return hash.Hash, nil
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
