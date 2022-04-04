package albumTestDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/test"
	"testing"
)

var tester structsTesters.HandlerTester

func init() {
	const testDataBaseSize = 20
	localStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})
	repo := *localStorage.GetAlbumRepo()
	storage.InitStorage(testDataBaseSize, &localStorage)
	tester = structsTesters.HandlerTester{}
	tester, _ = tester.SetTestingHandler(albumDeliveryHttp.Handler)
	testingHandler, _ := tester.GetTestingHandler()
	useCase, _ := testingHandler.GetUseCase()
	useCase, _ = useCase.SetRepo(repo, domain.AlbumMutex)
}

type TestDomainCreator struct{}

func (creator TestDomainCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Album{
		Id:             5, // for create test id will set automatically
		Title:          "testedAlbum",
		ArtistId:       7,
		CountLikes:     500,
		CountListening: 10000,
		Date:           0,
		CoverId:        5000,
	}
}

//const albumDeliveryHttpTestUrl = "/album/delivery/http/album_handler_test_"
//const url = router.Proto + router.Host + albumDeliveryHttpTestUrl

// ----------------------------------------------------------------------
func TestGet(t *testing.T) {
	tester.Get(t, domain.AlbumMutex)
}

func TestGetAll(t *testing.T) {
	tester.GetAll(t, domain.AlbumMutex)
}

func TestCreate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Create(t, creator, domain.AlbumMutex)
}

func TestDelete(t *testing.T) {
	tester.Delete(t, domain.AlbumMutex)
}

func TestUpdate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Update(t, creator, domain.AlbumMutex)
}

func TestPopular(t *testing.T) {
	tester.GetPopular(t, domain.AlbumMutex)
}
