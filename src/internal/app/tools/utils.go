package tools

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"math/rand"
	"reflect"
	"sync"
)

func Check(err error, id uint64, lastId uint64, errorString string) error {
	if err != nil {
		return err
	}
	if id > lastId {
		return errors.New(errorString)
	}
	return nil
}

func RandomRune() string {
	return string('a' + rune(rand.Intn('z'-'a'+1)))
}

func RandomWord(maxLen int) string {
	word := ""
	for i := 0; i < maxLen; i++ {
		word += RandomRune()
	}
	return word
}

// ---------------------------------------------------------
func CreateDataTransfer(domainType reflect.Type, dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.DataTransfer, error) {
	albumDomainType := reflect.TypeOf(domain.Album{})
	artistDomainType := reflect.TypeOf(domain.Artist{})
	trackDomainType := reflect.TypeOf(domain.Track{})

	if domainType == albumDomainType {
		artistId := dom.(domain.Album).ArtistId
		artistInCurrentAlbum, _ := artistUseCase.UseCase.GetById(artistId, mutex)
		return dom.CastDomainToDataTransferObject(*artistInCurrentAlbum)
	} else if domainType == artistDomainType {
		return dom.CastDomainToDataTransferObject(nil)
	} else if domainType == trackDomainType {
		artistId := dom.(domain.Track).ArtistId
		artistInCurrentTrack, _ := artistUseCase.UseCase.GetById(artistId, mutex)
		return dom.CastDomainToDataTransferObject(*artistInCurrentTrack)
	} else {
		return nil, errors.New(constants.BadType)
	}
}
