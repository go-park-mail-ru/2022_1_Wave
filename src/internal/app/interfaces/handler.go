package utilsInterfaces

import (
	"github.com/labstack/echo/v4"
	"reflect"
	"sync"
)

type HandlerInterface interface {
	GetAll(ctx echo.Context, mutex *sync.RWMutex) error
	Create(ctx echo.Context, mutex *sync.RWMutex) (HandlerInterface, error)
	Update(ctx echo.Context, mutex *sync.RWMutex) (HandlerInterface, error)
	Get(ctx echo.Context, mutex *sync.RWMutex) error
	Delete(ctx echo.Context, mutex *sync.RWMutex) (HandlerInterface, error)
	GetPopular(ctx echo.Context, mutex *sync.RWMutex) error
	GetModel() (reflect.Type, error)
	GetUseCase() (UseCaseInterface, error)
	SetModel(model reflect.Type) (HandlerInterface, error)
	SetUseCase(useCase UseCaseInterface, mutex *sync.RWMutex) (HandlerInterface, error)
}
