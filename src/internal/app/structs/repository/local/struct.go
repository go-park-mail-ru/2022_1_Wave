package structRepoLocal

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"math"
	"reflect"
	"sort"
	"sync"
)

type Repo struct {
	Domains []utilsInterfaces.Domain `json:"domain"`
}

// ----------------------------------------------------------------------
func (repo Repo) Insert(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()

	id, err := repo.GetSize(mutex)
	if err != nil {
		return nil, err
	}

	dom, err = dom.SetId(id)
	if err != nil {
		return nil, err
	}

	repo.Domains = append(repo.Domains, dom)
	return repo, nil
}

func (repo Repo) Update(domain utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	domainFromDB, err := repo.SelectByID(domain.GetId(), mutex)

	mutex.Lock()
	defer mutex.Unlock()

	if err != nil {
		return repo, err
	}

	repo.Domains[domainFromDB.GetId()] = domain

	return repo, nil
}

func (repo Repo) Delete(id uint64, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if id >= uint64(len(repo.Domains)) {
		return repo, errors.New(constants.IndexOutOfRange)
	}
	repo.Domains = append(repo.Domains[:id], repo.Domains[id+1:]...)

	for idx, domain := range repo.Domains {
		var err error
		repo.Domains[idx], err = domain.SetId(uint64(idx))
		if err != nil {
			return repo, err
		}
	}
	return repo, nil
}

func (repo Repo) SelectByID(id uint64, mutex *sync.RWMutex) (utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	return repo.Domains[id], nil
}

func (repo Repo) GetAll(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return repo.Domains, nil
}

func (repo Repo) GetPopular(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	var domainsPtr = make([]*utilsInterfaces.Domain, len(repo.Domains))
	for i := 0; i < len(repo.Domains); i++ {
		domainsPtr[i] = &repo.Domains[i]
	}

	sort.SliceStable(domainsPtr, func(i int, j int) bool {
		domain1 := *(domainsPtr[i])
		domain2 := *(domainsPtr[j])
		return domain1.GetCountListening() > domain2.GetCountListening()
	})

	topChart := make([]utilsInterfaces.Domain, uint64(math.Min(constants.Top, float64(len(repo.Domains)))))
	for i := 0; i < len(topChart); i++ {
		topChart[i] = *domainsPtr[i]
	}

	return topChart, nil
}

func (repo Repo) GetLastId(mutex *sync.RWMutex) (uint64, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if len(repo.Domains)-1 < 0 {
		return constants.NullId, errors.New(constants.ErrorDbIsEmpty)
	}

	return uint64(len(repo.Domains) - 1), nil
}

func (repo Repo) GetType(mutex *sync.RWMutex) reflect.Type {
	mutex.RLock()
	defer mutex.RUnlock()
	return reflect.TypeOf(repo)
}

// todo костыль
func (repo Repo) GetTracksFromAlbum(albumid uint64, mutex *sync.RWMutex) (interface{}, error) {
	return nil, nil
}

// todo пока кастыль, так как не успеваем
func (repo Repo) GetAlbumsFromArtist(artistId uint64, mutex *sync.RWMutex) (interface{}, error) {
	return nil, nil
}

func (repo Repo) GetSize(mutex *sync.RWMutex) (uint64, error) {
	//mutex.RLock()
	//defer mutex.RUnlock()
	return uint64(len(repo.Domains)), nil
}

// ----------------------------------------------------------------------
