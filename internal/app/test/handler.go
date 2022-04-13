package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	dataTransferCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/dataTransfer"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type HandlerTester struct {
	handler utilsInterfaces.HandlerInterface
}

func (tester HandlerTester) Get(t *testing.T) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	var dataTransferType reflect.Type

	model, err := tester.handler.GetModel()
	require.NoError(t, err)

	dataTransferType, err = dataTransferCreator.GetDataTransferTypeByDomainType(model)
	require.NoError(t, err)

	cases, err := PrepareArrayCases(useCase)

	for _, testCase := range cases {
		url := router.Proto + router.Host + "/" + router.Get + fmt.Sprint(testCase.Id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(url)
		ctx.SetParamNames("id")
		ctx.SetParamValues(fmt.Sprint(testCase.Id))

		require.NoError(t, tester.handler.Get(ctx))

		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var result webUtils.Success
		err = json.Unmarshal(body, &result)
		require.NoError(t, err)

		data := result.Result.(interface{})
		dataTransfer, err := dataTransferCreator.CreateDataTransferFromInterface(dataTransferType, data)
		require.NoError(t, err)

		require.Equal(t, testCase.Status, rec.Code)
		require.Equal(t, testCase.Data, dataTransfer)
	}

}

func (tester HandlerTester) GetAll(t *testing.T) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	repo, err := useCase.GetRepo()
	require.NoError(t, err)

	var dataTransferType reflect.Type
	model, err := tester.handler.GetModel()
	require.NoError(t, err)

	dataTransferType, err = dataTransferCreator.GetDataTransferTypeByDomainType(model)

	require.NoError(t, err)

	url := router.Proto + router.Host + "/" + router.Get
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(url)

	require.NoError(t, tester.handler.GetAll(ctx))

	testCases := PrepareManyCases(repo)

	require.Equal(t, testCases.Status, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	data := result.Result.([]interface{})
	ptr, err := dataTransferCreator.ToDataTransfers(dataTransferType, data)
	require.NoError(t, err)

	objects := *ptr

	for idx, testCase := range testCases.Data {
		require.Equal(t, testCase, objects[idx])
	}
}

func (tester HandlerTester) Create(t *testing.T, creator utilsInterfaces.TestDomainsCreator) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	sizeBefore, err := useCase.GetSize()
	require.NoError(t, err)

	sizeAfter := sizeBefore + 1

	testDomain := creator.PrepareOneTestDomain()

	url := router.Proto + router.Host + "/" + router.Create

	dataToSend, err := json.Marshal(testDomain)

	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataToSend))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(url)

	changedHandler, err := tester.handler.Create(ctx)
	require.NoError(t, err)

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	require.Equal(t, testCase.Status, rec.Code)

	tester.handler = changedHandler

	useCase, err = tester.handler.GetUseCase()
	require.NoError(t, err)

	lastId, err := useCase.GetLastId()
	require.NoError(t, err)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")",
	}

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)
	require.Equal(t, expected, result)

	useCase, err = tester.handler.GetUseCase()
	require.NoError(t, err)
	resultSize, err := useCase.GetSize()
	require.NoError(t, err)
	require.Equal(t, sizeAfter, resultSize)

}

func (tester HandlerTester) Delete(t *testing.T, idToDelete uint64) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	sizeBefore, err := useCase.GetSize()
	require.NoError(t, err)
	sizeBefore++
	sizeAfter := sizeBefore - 1

	domainToDelete, err := useCase.GetById(idToDelete)
	require.NoError(t, err)

	id := domainToDelete.GetId()
	require.Equal(t, id, idToDelete)

	url := router.Proto + router.Host + "/" + router.Delete + fmt.Sprint(id)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(url)
	ctx.SetParamNames(constants.FieldId)
	ctx.SetParamValues(fmt.Sprint(id))

	changedHandler, err := tester.handler.Delete(ctx)

	tester.handler = changedHandler

	require.NoError(t, err)

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	require.Equal(t, testCase.Status, rec.Code)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")",
	}

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	require.Equal(t, expected, result)

	useCase, err = tester.handler.GetUseCase()
	require.NoError(t, err)
	resultSize, err := useCase.GetSize()
	require.NoError(t, err)
	resultSize++

	require.Equal(t, sizeAfter, resultSize)
}

func (tester HandlerTester) Update(t *testing.T, creator utilsInterfaces.TestDomainsCreator) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	testDomain := creator.PrepareOneTestDomain()

	id := testDomain.GetId()

	url := router.Proto + router.Host + "/" + router.Update + fmt.Sprint(id)
	dataToSend, err := json.Marshal(testDomain)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(dataToSend))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(url)

	changedHandler, err := tester.handler.Update(ctx)
	require.NoError(t, err)
	tester.handler = changedHandler

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	require.Equal(t, testCase.Status, rec.Code)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")",
	}

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)
	require.Equal(t, expected, result)

	useCase, err = tester.handler.GetUseCase()
	require.NoError(t, err)
	dom, err := useCase.GetById(id)

	require.Equal(t, testDomain, dom)
}

func (tester HandlerTester) GetPopular(t *testing.T) {
	e := echo.New()

	useCase, err := tester.handler.GetUseCase()
	require.NoError(t, err)

	var dataTransferType reflect.Type
	model, err := tester.handler.GetModel()
	require.NoError(t, err)

	dataTransferType, err = dataTransferCreator.GetDataTransferTypeByDomainType(model)
	require.NoError(t, err)

	url := router.Proto + router.Host + "/" + router.Get
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath(url)

	require.NoError(t, tester.handler.GetPopular(ctx))

	testCase := PreparePopularCases(useCase)

	require.Equal(t, testCase.Status, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	data := result.Result.([]interface{})
	ptr, err := dataTransferCreator.ToDataTransfers(dataTransferType, data)
	require.NoError(t, err)

	objects := *ptr

	for idx, testCase := range testCase.Data {
		require.Equal(t, objects[idx], testCase)
	}
}

func (tester HandlerTester) SetHandler(handler utilsInterfaces.HandlerInterface) (HandlerTester, error) {
	tester.handler = handler
	return tester, nil
}

func (tester HandlerTester) GetHandler() (utilsInterfaces.HandlerInterface, error) {
	return tester.handler, nil
}
