package trackTestDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/test"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	"testing"
)

var tester structsTesters.HandlerTester

func init() {
	const testDataBaseSize = 20

	localStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})
	repo := *localStorage.GetTrackRepo()
	_ = storage.InitStorage(testDataBaseSize, &localStorage)

	tester = structsTesters.HandlerTester{}
	tester, _ = tester.SetTestingHandler(trackDeliveryHttp.Handler)
	testingHandler, _ := tester.GetTestingHandler()

	useCase, _ := testingHandler.GetUseCase()
	useCase, _ = useCase.SetRepo(repo, domain.TrackMutex)
}

type TestDomainCreator struct{}

func (creator TestDomainCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Track{
		Id:             7,
		AlbumId:        5,
		ArtistId:       3,
		Title:          "testTrack",
		Duration:       300,
		Mp4:            "someMp4 source",
		CoverId:        7,
		CountLikes:     5050,
		CountListening: 228,
	}
}

//const trackDeliveryHttpTestUrl = "/track/delivery/http/track_handler_test_"
//const url = router.Proto + router.Host + trackDeliveryHttpTestUrl

// ----------------------------------------------------------------------
func TestGet(t *testing.T) {
	tester.Get(t, domain.TrackMutex)
}

func TestGetAll(t *testing.T) {
	tester.GetAll(t, domain.TrackMutex)
}

func TestCreate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Create(t, creator, domain.TrackMutex)
}

func TestDelete(t *testing.T) {
	tester.Delete(t, domain.TrackMutex)
}

func TestUpdate(t *testing.T) {
	creator := TestDomainCreator{}
	tester.Update(t, creator, domain.TrackMutex)
}

func TestPopular(t *testing.T) {
	tester.GetPopular(t, domain.TrackMutex)
}
