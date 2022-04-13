package structsDeliveryHttp

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/dataTransfer"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type Handler struct {
	useCase utilsInterfaces.UseCaseInterface
	model   reflect.Type
}

func (h Handler) GetAll(ctx echo.Context) error {

	domains, err := h.useCase.GetAll()

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = []utilsInterfaces.Domain{}
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(domains))

	for i, dom := range domains {
		dataTransfer, err := dataTransferCreator.CreateDataTransfer(dom)
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

func (h Handler) Create(ctx echo.Context) (utilsInterfaces.HandlerInterface, error) {
	domainType := h.model

	objectToCreate, err := readPostPutRequest(ctx, domainType)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Create(objectToCreate)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.useCase.GetLastId()
	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
}

func (h Handler) Update(ctx echo.Context) (utilsInterfaces.HandlerInterface, error) {
	domainType := h.model

	objectToUpdate, err := readPostPutRequest(ctx, domainType)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Update(objectToUpdate)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := objectToUpdate.GetId()
	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

func (h Handler) Get(ctx echo.Context) error {
	id, err := ReadGetDeleteRequest(ctx)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dom, err := h.useCase.GetById(uint64(id))

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dataTransfer, err := dataTransferCreator.CreateDataTransfer(dom)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: dataTransfer})
}

func (h Handler) Delete(ctx echo.Context) (utilsInterfaces.HandlerInterface, error) {
	id, err := ReadGetDeleteRequest(ctx)

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	h.useCase, err = h.useCase.Delete(uint64(id))

	if err != nil {
		return h, webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return h, ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

func (h Handler) GetPopular(ctx echo.Context) error {
	popular, err := h.useCase.GetPopular()
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(popular))

	for i, pop := range popular {
		dataTransfer, err := dataTransferCreator.CreateDataTransfer(pop)
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

func (h Handler) SetUseCase(useCase utilsInterfaces.UseCaseInterface) (utilsInterfaces.HandlerInterface, error) {
	h.useCase = useCase
	return h, nil
}
