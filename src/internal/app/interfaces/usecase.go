package utilsInterfaces

import (
	"reflect"
	"sync"
)

type UseCaseInterface interface {
	GetAll(mutex *sync.RWMutex) (*[]Domain, error)
	GetLastId(mutex *sync.RWMutex) (id uint64, err error)
	Create(*Domain, *sync.RWMutex) (UseCaseInterface, error)
	Update(Domain, *sync.RWMutex) (UseCaseInterface, error)
	Delete(id uint64, mutex *sync.RWMutex) (UseCaseInterface, error)
	GetById(id uint64, mutex *sync.RWMutex) (*Domain, error)
	GetPopular(*sync.RWMutex) (*[]Domain, error)
	GetType() reflect.Type
	SetRepo(repoInterface RepoInterface, mutex *sync.RWMutex) (UseCaseInterface, error)
	GetRepo() (RepoInterface, error)
	GetSize(mutex *sync.RWMutex) (uint64, error)
}
