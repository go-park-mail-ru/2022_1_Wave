package structRepoLocal

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"math"
	"reflect"
	"sort"
	"sync"
)

type Repo struct {
	Domains []utilsInterfaces.Domain `json:"domain"`
}

// ----------------------------------------------------------------------
func (storage Repo) Insert(domain *utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()
	storage.Domains = append(storage.Domains, *domain)
	return storage, nil
}

func (storage Repo) Update(domain *utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	domainFromDB, err := storage.SelectByID((*domain).GetId(), mutex)

	mutex.Lock()
	defer mutex.Unlock()

	if err != nil {
		return storage, err
	}
	*domainFromDB = *domain

	return storage, nil
}

func (storage Repo) Delete(id uint64, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if id >= uint64(len(storage.Domains)) {
		return storage, errors.New(constants.IndexOutOfRange)
	}
	storage.Domains = append(storage.Domains[:id], storage.Domains[id+1:]...)

	for idx, domain := range storage.Domains {
		var err error
		storage.Domains[idx], err = domain.SetId(uint64(idx))
		if err != nil {
			return storage, err
		}
	}
	return storage, nil
}

func (storage Repo) SelectByID(id uint64, mutex *sync.RWMutex) (*utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if id+1 > uint64(len(storage.Domains)) {
		return nil, errors.New(constants.IndexOutOfRange)
	}
	return &storage.Domains[id], nil
}

func (storage Repo) GetAll(mutex *sync.RWMutex) (*[]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return &storage.Domains, nil
}

func (storage Repo) GetPopular(mutex *sync.RWMutex) (*[]utilsInterfaces.Domain, error) {
	const top = 20
	mutex.RLock()
	defer mutex.RUnlock()

	var domainsPtr = make([]*utilsInterfaces.Domain, len(storage.Domains))
	for i := 0; i < len(storage.Domains); i++ {
		domainsPtr[i] = &storage.Domains[i]
	}

	sort.SliceStable(domainsPtr, func(i int, j int) bool {
		domain1 := *(domainsPtr[i])
		domain2 := *(domainsPtr[j])
		return domain1.GetCountListening() > domain2.GetCountListening()
	})

	topChart := make([]utilsInterfaces.Domain, uint64(math.Min(top, float64(len(storage.Domains)))))
	for i := 0; i < len(topChart); i++ {
		topChart[i] = *domainsPtr[i]
	}

	return &topChart, nil
}

func (storage Repo) GetLastId(mutex *sync.RWMutex) (uint64, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if len(storage.Domains)-1 < 0 {
		return constants.NullId, errors.New(constants.ErrorDbIsEmpty)
	}

	return uint64(len(storage.Domains) - 1), nil
}

func (storage Repo) GetType(mutex *sync.RWMutex) reflect.Type {
	mutex.RLock()
	defer mutex.RUnlock()
	return reflect.TypeOf(storage)
}

func (storage Repo) GetRefDomains(mutex *sync.RWMutex) []utilsInterfaces.Domain {
	mutex.RLock()
	defer mutex.RUnlock()
	return storage.Domains
}

// ----------------------------------------------------------------------
