package structsDeliveryHttp

import (
	"encoding/json"
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"sync"
)

type Handler struct {
	useCase utilsInterfaces.UseCaseInterface
	model   reflect.Type
}

func (h Handler) GetAll(w http.ResponseWriter, mutex *sync.RWMutex) {
	domainType := h.model

	domains, err := h.useCase.GetAll(mutex)
	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if *domains == nil {
		*domains = []utilsInterfaces.Domain{}
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(*domains))

	for i, dom := range *domains {
		dataTransfer, err := tools.CreateDataTransfer(domainType, dom, mutex)
		if err != nil {
			webUtils.WriteError(w, err, http.StatusBadRequest)
		}

		dataTransfers[i] = dataTransfer
	}

	_ = json.NewEncoder(w).Encode(webUtils.Success{
		Status: webUtils.OK,
		Result: dataTransfers})

}

func (h Handler) Create(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) utilsInterfaces.UseCaseInterface {
	domainType := h.model

	objectToCreate, err := readPostPutRequest(r, domainType)
	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	h.useCase, err = h.useCase.Create(objectToCreate, mutex)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	_, err = h.useCase.GetLastId(mutex)
	_ = json.NewEncoder(w).Encode(webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessCreated,
	})

	return h.useCase

}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) utilsInterfaces.UseCaseInterface {
	domainType := h.model

	objectToUpdate, err := readPostPutRequest(r, domainType)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	h.useCase, err = h.useCase.Update(objectToUpdate, mutex)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	id := objectToUpdate.GetId()
	_ = json.NewEncoder(w).Encode(webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")",
	})

	return h.useCase
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) {
	domainType := h.model

	id, err := readGetDeleteRequest(r)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	dom, err := h.useCase.GetById(uint64(id), mutex)
	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	dataTransfer, err := tools.CreateDataTransfer(domainType, *dom, mutex)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(webUtils.Success{
		Status: webUtils.OK,
		Result: dataTransfer})
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, mutex *sync.RWMutex) utilsInterfaces.UseCaseInterface {
	id, err := readGetDeleteRequest(r)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	h.useCase, err = h.useCase.Delete(uint64(id), mutex)

	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return h.useCase
	}

	_ = json.NewEncoder(w).Encode(webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")",
	})
	return h.useCase
}

func (h Handler) GetPopular(w http.ResponseWriter, mutex *sync.RWMutex) {
	popular, err := h.useCase.GetPopular(mutex)
	if err != nil {
		webUtils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	dataTransfers := make([]utilsInterfaces.DataTransfer, len(*popular))

	domainType := h.model

	for i, pop := range *popular {
		dataTransfer, err := tools.CreateDataTransfer(domainType, pop, mutex)
		if err != nil {
			webUtils.WriteError(w, err, http.StatusBadRequest)
		}

		dataTransfers[i] = dataTransfer

	}

	_ = json.NewEncoder(w).Encode(webUtils.Success{
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

func readPostPutRequest(r *http.Request, domainType reflect.Type) (utilsInterfaces.Domain, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var result interface{}

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	concreteDomain, errDueCast := tools.Creator.CreateDomainFromInterface(domainType, result)

	if errDueCast != nil {
		return nil, err
	}

	if err = concreteDomain.Check(); err != nil {
		return nil, err
	}

	object := concreteDomain

	return object, nil
}

func readGetDeleteRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[constants.FieldId])
	if err != nil {
		return constants.BadId, err
	}

	if id < 0 {
		return constants.BadId, errors.New(constants.IndexOutOfRange)
	}

	return id, nil
}
