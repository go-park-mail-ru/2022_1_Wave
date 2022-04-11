package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	structRepoPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/repository/postgresql"
	dataTransferCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/dataTransfer"
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

	for _, object := range objects {
		dataTransfer, _ := dataTransferCreator.CreateDataTransfer(object, mutex)
		cases.Data = append(cases.Data, dataTransfer)
	}

	cases.Status = http.StatusOK
	return cases
}

//func PrepareArrayCases(repo utilsInterfaces.RepoInterface, mutex *sync.RWMutex) []TestCase {
//	objects, _ := repo.GetAll(mutex)
//	cases := make([]TestCase, len(*objects))
//
//	fmt.Println(objects)
//	fmt.Println(reflect.TypeOf(objects))
//
//	domainType := reflect.TypeOf((*objects)[0])
//
//	for idx, object := range *objects {
//		fmt.Println(reflect.TypeOf(object))
//		dataTransfer, _ := dataTransferCreator.CreateDataTransfer(domainType, object, mutex)
//		cases[idx] = TestCase{
//			Id:     idx,
//			Status: http.StatusOK,
//			Data:   dataTransfer,
//		}
//	}
//	return cases
//}

func PrepareArrayCases(useCase utilsInterfaces.UseCaseInterface, mutex *sync.RWMutex) ([]TestCase, error) {
	objects, err := useCase.GetAll(mutex)

	if err != nil {
		return nil, err
	}
	cases := make([]TestCase, len(objects))

	repo, err := useCase.GetRepo()
	if err != nil {
		return nil, err
	}

	repoType := repo.GetType(mutex)
	postgresType := reflect.TypeOf(structRepoPostgres.Table{})

	isPostgres := repoType == postgresType

	for idx, obj := range objects {
		dataTransfer, _ := dataTransferCreator.CreateDataTransfer(obj, mutex)

		id := idx
		if isPostgres {
			id = idx + 1
		}

		cases[idx] = TestCase{
			Id:     id,
			Status: http.StatusOK,
			Data:   dataTransfer,
		}
	}
	return cases, nil
}

func PreparePopularCases(useCase utilsInterfaces.UseCaseInterface, mutex *sync.RWMutex) TestCases {
	cases := TestCases{}
	objects, _ := useCase.GetPopular(mutex)

	for _, object := range objects {
		dataTransfer, _ := dataTransferCreator.CreateDataTransfer(object, mutex)
		cases.Data = append(cases.Data, dataTransfer)
	}
	cases.Status = http.StatusOK
	return cases
}
