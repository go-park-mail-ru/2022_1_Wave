package utilsInterfaces

import (
	"reflect"
)

type UseCaseInterface interface {
	GetAll() ([]DataTransfer, error)
	GetLastId() (id int, err error)
	Create(Domain) error
	Update(Domain) error
	Delete(id int) error
	GetById(id int) (DataTransfer, error)
	GetType() reflect.Type
	//SetRepo(repoInterface RepoInterface) error
	//GetRepo() (RepoInterface, error)
	GetSize() (int, error)
}
