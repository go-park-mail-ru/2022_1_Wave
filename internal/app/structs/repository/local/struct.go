package structRepoLocal

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"math"
	"reflect"
	"sort"
)

type Repo struct {
	Domains []utilsInterfaces.Domain `json:"domain"`
}

// ----------------------------------------------------------------------
func (repo Repo) Insert(dom utilsInterfaces.Domain) (utilsInterfaces.RepoInterface, error) {
	id, err := repo.GetSize()
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

func (repo Repo) Update(domain utilsInterfaces.Domain) (utilsInterfaces.RepoInterface, error) {
	domainFromDB, err := repo.SelectByID(domain.GetId())

	if err != nil {
		return repo, err
	}

	repo.Domains[domainFromDB.GetId()] = domain

	return repo, nil
}

//todo убрать мьютексы

func (repo Repo) Delete(id uint64) (utilsInterfaces.RepoInterface, error) {
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

func (repo Repo) SelectByID(id uint64) (utilsInterfaces.Domain, error) {
	return repo.Domains[id], nil
}

func (repo Repo) GetAll() ([]utilsInterfaces.Domain, error) {
	return repo.Domains, nil
}

func (repo Repo) GetPopular() ([]utilsInterfaces.Domain, error) {
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

func (repo Repo) GetLastId() (uint64, error) {
	if len(repo.Domains)-1 < 0 {
		return constants.NullId, errors.New(constants.ErrorDbIsEmpty)
	}

	return uint64(len(repo.Domains) - 1), nil
}

func (repo Repo) GetType() reflect.Type {
	return reflect.TypeOf(repo)
}

// todo костыль
func (repo Repo) GetTracksFromAlbum(albumid uint64) (interface{}, error) {
	return nil, nil
}

// todo пока кастыль, так как не успеваем
func (repo Repo) GetAlbumsFromArtist(artistId uint64) (interface{}, error) {
	return nil, nil
}

// todo пока кастыль, так как не успеваем
func (repo Repo) GetPopularTracksFromArtist(artistId uint64) (interface{}, error) {
	return nil, nil
}

func (repo Repo) GetSize() (uint64, error) {
	//.RLock()
	//defer .RUnlock()
	return uint64(len(repo.Domains)), nil
}

// ----------------------------------------------------------------------
