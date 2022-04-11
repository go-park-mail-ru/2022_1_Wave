package structsUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"reflect"
	"sync"
)

type UseCase struct {
	repo utilsInterfaces.RepoInterface
}

func (useCase UseCase) GetAll(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	return useCase.repo.GetAll(mutex)
}

func (useCase UseCase) GetLastId(mutex *sync.RWMutex) (id uint64, err error) {
	return useCase.repo.GetLastId(mutex)
}

func (useCase UseCase) Create(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Insert(dom, mutex)
	return useCase, err
}

func (useCase UseCase) Update(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Update(dom, mutex)
	return useCase, err
}

func (useCase UseCase) Delete(id uint64, mutex *sync.RWMutex) (utilsInterfaces.UseCaseInterface, error) {
	var err error
	useCase.repo, err = useCase.repo.Delete(id, mutex)
	return useCase, err
}

func (useCase UseCase) GetById(id uint64, mutex *sync.RWMutex) (utilsInterfaces.Domain, error) {
	return useCase.repo.SelectByID(id, mutex)
}

func (useCase UseCase) GetPopular(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	return useCase.repo.GetPopular(mutex)
}

func (useCase UseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.Artist{})
}

func (useCase UseCase) GetRepo() (utilsInterfaces.RepoInterface, error) {
	return useCase.repo, nil
}

func (useCase UseCase) SetRepo(repo utilsInterfaces.RepoInterface, mutex *sync.RWMutex) (utilsInterfaces.UseCaseInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()
	useCase.repo = repo
	return useCase, nil
}

func (useCase UseCase) GetSize(mutex *sync.RWMutex) (uint64, error) {
	return useCase.repo.GetSize(mutex)
}
