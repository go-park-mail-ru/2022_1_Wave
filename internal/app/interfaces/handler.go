package utilsInterfaces

import (
	"github.com/labstack/echo/v4"
	"reflect"
)

type HandlerInterface interface {
	GetAll(ctx echo.Context) error
	Create(ctx echo.Context) (HandlerInterface, error)
	Update(ctx echo.Context) (HandlerInterface, error)
	Get(ctx echo.Context) error
	Delete(ctx echo.Context) (HandlerInterface, error)
	GetPopular(ctx echo.Context) error
	GetModel() (reflect.Type, error)
	GetUseCase() (UseCaseInterface, error)
	SetModel(model reflect.Type) (HandlerInterface, error)
	SetUseCase(useCase UseCaseInterface) (HandlerInterface, error)
}
