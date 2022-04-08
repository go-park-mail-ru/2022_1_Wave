package artistTestDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/test"
	"testing"
)

var tester structsTesters.HandlerTester

func init() {
	const testDataBaseSize = 20

	localStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})
	repo := *localStorage.GetArtistRepo()
	_ = storage.InitStorage(testDataBaseSize, &localStorage)

	tester = structsTesters.HandlerTester{}
	tester, _ = tester.SetTestingHandler(artistDeliveryHttp.Handler)
	testingHandler, _ := tester.GetTestingHandler()

	useCase, _ := testingHandler.GetUseCase()
	useCase, _ = useCase.SetRepo(repo, domain.ArtistMutex)

}

type TestDomainCreator struct{}

func (creator TestDomainCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Artist{
		Id:             5,
		Name:           "testArtist",
		CountFollowers: 1355,
		CountListening: 5000,
	}
}

//const albumDeliveryHttpTestUrl = "/album/delivery/http/album_handler_test_"
//const url = router.Proto + router.Host + albumDeliveryHttpTestUrl

// ----------------------------------------------------------------------
func TestGet(t *testing.T) {
	tester.Get(t, domain.ArtistMutex)
}

func TestGetAll(t *testing.T) {
	tester.GetAll(t, domain.ArtistMutex)
}

func TestCreate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Create(t, creator, domain.ArtistMutex)
}

func TestDelete(t *testing.T) {
	tester.Delete(t, domain.ArtistMutex)
}

func TestUpdate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Update(t, creator, domain.ArtistMutex)
}

func TestPopular(t *testing.T) {
	tester.GetPopular(t, domain.ArtistMutex)
}
