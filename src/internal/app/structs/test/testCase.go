package structsTesters

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools"
	"net/http"
	"reflect"
	"sync"
)

type TestCase struct {
	Id     int
	Status int
	Data   utilsInterfaces.DataTransfer
}

type TestCases struct {
	Data   []utilsInterfaces.DataTransfer
	Status int
}

type OperationTestCase struct {
	Status int
}

func PrepareManyCases(repo utilsInterfaces.RepoInterface, mutex *sync.RWMutex) TestCases {
	cases := TestCases{}
	objects, _ := repo.GetAll(mutex)

	domainType := reflect.TypeOf((*objects)[0])

	for _, object := range *objects {
		dataTransfer, _ := tools.CreateDataTransfer(domainType, object, mutex)
		cases.Data = append(cases.Data, dataTransfer)
	}

	cases.Status = http.StatusOK
	return cases
}

func PrepareArrayCases(repo utilsInterfaces.RepoInterface, mutex *sync.RWMutex) []TestCase {
	objects, _ := repo.GetAll(mutex)
	cases := make([]TestCase, len(*objects))

	domainType := reflect.TypeOf((*objects)[0])

	for idx, object := range *objects {
		dataTransfer, _ := tools.CreateDataTransfer(domainType, object, mutex)
		cases[idx] = TestCase{
			Id:     idx,
			Status: http.StatusOK,
			Data:   dataTransfer,
		}
	}
	return cases
}

func PreparePopularCases(repo utilsInterfaces.RepoInterface, mutex *sync.RWMutex) TestCases {
	cases := TestCases{}
	objects, _ := repo.GetPopular(mutex)

	domainType := reflect.TypeOf((*objects)[0])

	for _, object := range *objects {
		dataTransfer, _ := tools.CreateDataTransfer(domainType, object, mutex)
		cases.Data = append(cases.Data, dataTransfer)
	}
	cases.Status = http.StatusOK
	return cases
}
