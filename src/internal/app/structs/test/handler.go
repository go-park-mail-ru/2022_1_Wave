package structsTesters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
)

type HandlerTester struct {
	handler utilsInterfaces.HandlerInterface
}

func (tester HandlerTester) Get(t *testing.T, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()
	repo, _ := useCase.GetRepo()
	var dataTransferType reflect.Type
	model, _ := tester.handler.GetModel()
	dataTransferType, _ = tools.Converter.GetDataTransferTypeByDomainType(model)

	cases := PrepareArrayCases(repo, mutex)

	for caseNum, item := range cases {
		url := router.Proto + router.Host + "/" + router.Get + fmt.Sprint(item.Id)
		req := httptest.NewRequest("GET", url, nil)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": fmt.Sprint(item.Id),
		}

		req = mux.SetURLVars(req, vars)
		tester.handler.Get(w, req, mutex)

		if w.Code != item.Status {
			t.Fatalf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.Status)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		var result webUtils.Success
		_ = json.Unmarshal(body, &result)
		data := result.Result.(interface{})

		dataTransfer, _ := tools.Creator.CreateDataTransferFromInterface(dataTransferType, data)

		if dataTransfer != item.Data.(utilsInterfaces.DataTransfer) {
			t.Fatalf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, resp, item.Data)
		}
	}

}

func (tester HandlerTester) GetAll(t *testing.T, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()
	repo, _ := useCase.GetRepo()
	var dataTransferType reflect.Type
	model, _ := tester.handler.GetModel()
	dataTransferType, _ = tools.Converter.GetDataTransferTypeByDomainType(model)

	w := httptest.NewRecorder()

	tester.handler.GetAll(w, mutex)

	testCases := PrepareManyCases(repo, mutex)

	if w.Code != testCases.Status {
		t.Fatalf("wrong StatusCode: got %d, expected %d", w.Code, testCases.Status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result webUtils.Success
	_ = json.Unmarshal(body, &result)
	data := result.Result.([]interface{})
	ptr, err := tools.Converter.ToDataTransfers(dataTransferType, data)

	objects := *ptr

	if err != nil {
		t.Fatalf("error to casting: %v", err)
	}

	for idx, testCase := range testCases.Data {
		if objects[idx] != testCase {
			t.Fatalf("wrong Response: got %+v, expected %+v", objects[idx], testCase)
		}
	}
}

func (tester HandlerTester) Create(t *testing.T, creator utilsInterfaces.TestDomainsCreator, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()

	sizeBefore, _ := useCase.GetLastId(mutex)
	sizeBefore++
	sizeAfter := sizeBefore + 1

	testDomain := creator.PrepareOneTestDomain()

	url := router.Proto + router.Host + "/" + router.Create

	dataToSend, _ := json.Marshal(testDomain)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	useCase = tester.handler.Create(w, req, mutex)

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	if w.Code != testCase.Status {
		t.Fatalf("wrong StatusCode: got %d, expected %d", w.Code, testCase.Status)
	}

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessCreated,
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result webUtils.Success
	_ = json.Unmarshal(body, &result)
	if result != expected {
		t.Fatalf("wrong Response: got %+v, expected %+v",
			result, expected)
	}

	resultSize, _ := useCase.GetLastId(mutex)
	resultSize++

	require.Equal(t, sizeAfter, resultSize)

}

func (tester HandlerTester) Delete(t *testing.T, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()

	const idToDelete = uint64(0)

	sizeBefore, _ := useCase.GetLastId(mutex)
	sizeBefore++
	sizeAfter := sizeBefore - 1

	domainToDelete, _ := useCase.GetById(idToDelete, mutex)
	id := (*domainToDelete).GetId()

	require.Equal(t, id, idToDelete)

	url := router.Proto + router.Host + "/" + router.Delete + fmt.Sprint(id)
	req := httptest.NewRequest("DELETE", url, nil)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(id),
	}

	req = mux.SetURLVars(req, vars)
	useCase = tester.handler.Delete(w, req, mutex)

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	if w.Code != testCase.Status {
		t.Fatalf("wrong StatusCode: got %d, expected %d", w.Code, testCase.Status)
	}

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")",
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result webUtils.Success
	_ = json.Unmarshal(body, &result)
	if result != expected {
		t.Fatalf("wrong Response: got %+v, expected %+v", result, expected)
	}
	resultSize, _ := useCase.GetLastId(mutex)
	resultSize++

	require.Equal(t, sizeAfter, resultSize)
}

func (tester HandlerTester) Update(t *testing.T, creator utilsInterfaces.TestDomainsCreator, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()

	testDomain := creator.PrepareOneTestDomain()

	id := testDomain.GetId()

	url := router.Proto + router.Host + "/" + router.Update + fmt.Sprint(id)
	dataToSend, _ := json.Marshal(testDomain)
	req := httptest.NewRequest("PUT", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()

	useCase = tester.handler.Update(w, req, mutex)

	testCase := OperationTestCase{
		Status: http.StatusOK,
	}

	if w.Code != testCase.Status {
		t.Fatalf("wrong StatusCode: got %d, expected %d", w.Code, testCase.Status)
	}

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")",
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result webUtils.Success
	_ = json.Unmarshal(body, &result)
	if result != expected {
		t.Fatalf("wrong Response: got %+v, expected %+v",
			result, expected)
	}

	dom, _ := useCase.GetById(id, mutex)
	require.Equal(t, testDomain, *dom)
}

func (tester HandlerTester) GetPopular(t *testing.T, mutex *sync.RWMutex) {
	useCase, _ := tester.handler.GetUseCase()
	repo, _ := useCase.GetRepo()
	var dataTransferType reflect.Type
	model, _ := tester.handler.GetModel()
	dataTransferType, _ = tools.Converter.GetDataTransferTypeByDomainType(model)

	w := httptest.NewRecorder()

	tester.handler.GetPopular(w, mutex)

	testCase := PreparePopularCases(repo, mutex)

	if w.Code != testCase.Status {
		t.Fatalf("wrong StatusCode: got %d, expected %d", w.Code, testCase.Status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result webUtils.Success

	_ = json.Unmarshal(body, &result)

	data := result.Result.([]interface{})
	ptr, err := tools.Converter.ToDataTransfers(dataTransferType, data)

	if err != nil {
		t.Fatalf("error to casting: %v", err)
	}

	objects := *ptr

	for idx, testCase := range testCase.Data {
		if objects[idx] != testCase {
			t.Fatalf("wrong Response: got %+v, expected %+v", objects[idx], testCase)
		}
	}
}

func (tester HandlerTester) SetTestingHandler(handler utilsInterfaces.HandlerInterface) (HandlerTester, error) {
	tester.handler = handler
	return tester, nil
}

func (tester HandlerTester) GetTestingHandler() (utilsInterfaces.HandlerInterface, error) {
	return tester.handler, nil
}
