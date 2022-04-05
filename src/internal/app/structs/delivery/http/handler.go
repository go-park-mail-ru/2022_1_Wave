package structsDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strconv"
	"sync"
)

type Handler struct {
	useCase utilsInterfaces.UseCaseInterface
	model   reflect.Type
}

func (h Handler) GetAll(ctx echo.Context, mutex *sync.RWMutex) error {
	domainType := h.model

	domains, err := h.useCase.GetAll(mutex)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if *domains == nil {
		*domains = []utilsInterfaces.Domain{}
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(*domains))

	for i, dom := range *domains {
		dataTransfer, err := tools.CreateDataTransfer(domainType, dom, mutex)
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}

		dataTransfers[i] = dataTransfer
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: dataTransfers})
}

func (h Handler) Create(ctx echo.Context, mutex *sync.RWMutex) (utilsInterfaces.HandlerInterface, error) {
	domainType := h.model

	objectToCreate, err := readPostPutRequest(ctx, domainType)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Create(&objectToCreate, mutex)

	fmt.Println(objectToCreate)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessCreated + "(" + fmt.Sprint(objectToCreate.GetId()) + ")"})
}

func (h Handler) Update(ctx echo.Context, mutex *sync.RWMutex) (utilsInterfaces.HandlerInterface, error) {
	domainType := h.model

	objectToUpdate, err := readPostPutRequest(ctx, domainType)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Update(objectToUpdate, mutex)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := objectToUpdate.GetId()
	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

func (h Handler) Get(ctx echo.Context, mutex *sync.RWMutex) error {
	domainType := h.model

	id, err := readGetDeleteRequest(ctx)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dom, err := h.useCase.GetById(uint64(id), mutex)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dataTransfer, err := tools.CreateDataTransfer(domainType, *dom, mutex)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: dataTransfer})
}

func (h Handler) Delete(ctx echo.Context, mutex *sync.RWMutex) (utilsInterfaces.HandlerInterface, error) {
	id, err := readGetDeleteRequest(ctx)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Delete(uint64(id), mutex)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

func (h Handler) GetPopular(ctx echo.Context, mutex *sync.RWMutex) error {
	popular, err := h.useCase.GetPopular(mutex)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(*popular))

	domainType := h.model

	for i, pop := range *popular {
		dataTransfer, err := tools.CreateDataTransfer(domainType, pop, mutex)
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}

		dataTransfers[i] = dataTransfer

	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: dataTransfers})
}

func (h Handler) GetModel() (reflect.Type, error) {
	return h.model, nil
}

func (h Handler) GetUseCase() (utilsInterfaces.UseCaseInterface, error) {
	return h.useCase, nil
}

func (h Handler) SetModel(model reflect.Type) (utilsInterfaces.HandlerInterface, error) {
	h.model = model
	return h, nil
}

func (h Handler) SetUseCase(useCase utilsInterfaces.UseCaseInterface, mutex *sync.RWMutex) (utilsInterfaces.HandlerInterface, error) {
	h.useCase = useCase
	return h, nil
}

func readPostPutRequest(ctx echo.Context, domainType reflect.Type) (utilsInterfaces.Domain, error) {
	var result interface{}

	if err := ctx.Bind(&result); err != nil {
		return nil, err
	}

	concreteDomain, errDueCast := tools.Creator.CreateDomainFromInterface(domainType, result)

	if errDueCast != nil {
		return nil, errDueCast
	}

	if err := concreteDomain.Check(); err != nil {
		return nil, err
	}

	object := concreteDomain

	return object, nil
}

func readGetDeleteRequest(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return constants.BadId, err
	}

	if id < 0 {
		return constants.BadId, errors.New(constants.IndexOutOfRange)
	}

	return id, nil
}
