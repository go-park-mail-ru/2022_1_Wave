package LinkerGrpc

import (
	"context"
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

func (useCase LinkerGrpc) Get(ctx context.Context, url *linkerProto.UrlWrapper) (*linkerProto.HashWrapper, error) {
	hash, err := (*useCase.LinkerRepo).Get(url.Url)
	if err != nil {
		return nil, err
	}
	return &linkerProto.HashWrapper{Hash: hash}, nil
}

func (useCase LinkerGrpc) Create(ctx context.Context, data *linkerProto.UrlWrapper) (*linkerProto.HashWrapper, error) {
	hash, err := (*useCase.LinkerRepo).Create(data.Url)
	return &linkerProto.HashWrapper{Hash: hash}, err
}
