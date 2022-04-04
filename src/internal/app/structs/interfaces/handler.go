package utilsInterfaces

import (
	"net/http"
	"reflect"
	"sync"
)

type HandlerInterface interface {
	GetAll(w http.ResponseWriter, mutex *sync.RWMutex)
	Create(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) UseCaseInterface
	Update(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) UseCaseInterface
	Get(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex)
	Delete(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) UseCaseInterface
	GetPopular(w http.ResponseWriter, mutex *sync.RWMutex)
	GetModel() (reflect.Type, error)
	GetUseCase() (UseCaseInterface, error)
	SetModel(model reflect.Type) (HandlerInterface, error)
	SetUseCase(useCase UseCaseInterface, mutex *sync.RWMutex) (HandlerInterface, error)
}
