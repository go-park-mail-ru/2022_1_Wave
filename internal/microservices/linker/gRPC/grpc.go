package LinkerGrpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/linkerProto"
)

type LinkerGrpc struct {
	LinkerRepo *domain.LinkerRepo
	linkerProto.UnimplementedLinkerUseCaseServer
}

func MakeLinkerGrpc(linkerRepo domain.LinkerRepo) LinkerGrpc {
	return LinkerGrpc{LinkerRepo: &linkerRepo}
}

func (useCase LinkerGrpc) Get(ctx context.Context, hash *linkerProto.HashWrapper) (*linkerProto.UrlWrapper, error) {
	url, err := (*useCase.LinkerRepo).Get(hash.Hash)
	if err != nil {
		return nil, err
	}
	return &linkerProto.UrlWrapper{Url: url}, nil
}

func (useCase LinkerGrpc) Create(ctx context.Context, data *linkerProto.UrlWrapper) (*linkerProto.HashWrapper, error) {
	hash, err := (*useCase.LinkerRepo).Create(data.Url)
	return &linkerProto.HashWrapper{Hash: hash}, err
}

func (useCase LinkerGrpc) Count(ctx context.Context, hash *linkerProto.HashWrapper) (*linkerProto.CountResponse, error) {
	fmt.Println("in grpc count")
	count, err := (*useCase.LinkerRepo).Count(hash.Hash)
	if err != nil {
		return nil, err
	}
	return &linkerProto.CountResponse{Result: count}, nil
}
