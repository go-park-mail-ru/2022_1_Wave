package structsUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"reflect"
)

type UseCase struct {
	repo utilsInterfaces.RepoInterface
}

func (useCase UseCase) GetAll() ([]utilsInterfaces.Domain, error) {
	return useCase.repo.GetAll()
}

func (useCase UseCase) GetLastId() (id uint64, err error) {
	return useCase.repo.GetLastId()
}

func (useCase UseCase) Create(dom utilsInterfaces.Domain) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Insert(dom)
	return useCase, err
}

func (useCase UseCase) Update(dom utilsInterfaces.Domain) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Update(dom)
	return useCase, err
}

func (useCase UseCase) Delete(id uint64) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Delete(id)
	return useCase, err
}

func (useCase UseCase) GetById(id uint64) (utilsInterfaces.Domain, error) {
	return useCase.repo.SelectByID(id)
}

func (useCase UseCase) GetPopular() ([]utilsInterfaces.Domain, error) {
	return useCase.repo.GetPopular()
}

func (useCase UseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.Artist{})
}

func (useCase UseCase) GetRepo() (utilsInterfaces.RepoInterface, error) {
	return useCase.repo, nil
}

func (useCase UseCase) GetTracksFromAlbum(albumId uint64) (interface{}, error) {
	return useCase.repo.GetTracksFromAlbum(albumId)
}

func (useCase UseCase) GetAlbumsFromArtist(artist uint64) (interface{}, error) {
	return useCase.repo.GetAlbumsFromArtist(artist)
}

func (useCase UseCase) GetPopularTracksFromArtist(artistId uint64) (interface{}, error) {
	return useCase.repo.GetPopularTracksFromArtist(artistId)
}

func (useCase UseCase) SetRepo(repo utilsInterfaces.RepoInterface) (utilsInterfaces.UseCaseInterface, error) {
	useCase.repo = repo
	return useCase, nil
}

func (useCase UseCase) GetSize() (uint64, error) {
	return useCase.repo.GetSize()
}
