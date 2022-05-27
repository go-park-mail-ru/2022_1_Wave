package LinkerUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
)

type LinkerUseCase interface {
	Get(hash string) (string, error)
	Create(url string) (string, error)
}

type linkerUseCase struct {
	linkerAgent domain.LinkerAgent
}

func NewLinkerUseCase(linkerAgent domain.LinkerRepo) *linkerUseCase {
	return &linkerUseCase{
		linkerAgent: linkerAgent,
	}
}

func (useCase linkerUseCase) Create(url string) (string, error) {
	returnedHash, err := useCase.linkerAgent.Create(url)
	return returnedHash, err
}

func (useCase linkerUseCase) Get(hash string) (string, error) {
	url, err := useCase.linkerAgent.Get(hash)
	return url, err
}
