package LinkerUseCase

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"net/url"
	"strings"
)

type LinkerUseCase interface {
	Get(hash string) (string, error)
	Create(url string) (string, error)
	Count(hash string) (int64, error)
}

type linkerUseCase struct {
	linkerAgent domain.LinkerAgent
}

func NewLinkerUseCase(linkerAgent domain.LinkerAgent) *linkerUseCase {
	return &linkerUseCase{
		linkerAgent: linkerAgent,
	}
}

const waveMusicHost = "wave-music.online"

func linkToFormat(link string) (*url.URL, error) {
	u, err := url.ParseRequestURI(link)
	if err != nil {
		return nil, err
	}

	u.Scheme = "https"

	parsedLink := u.String()

	if !govalidator.IsURL(parsedLink) {
		return nil, errors.New("invalid link: " + link)
	}

	delim := "://"
	wwwPart := "www."

	wwwBegin := strings.Index(parsedLink, delim+wwwPart)
	if wwwBegin != -1 {
		wwwEnd := wwwBegin + len(delim+wwwPart)
		parsedLink = parsedLink[:wwwBegin+len(delim)] + parsedLink[wwwEnd:]
	}

	fmt.Println("parsed link=", parsedLink)

	currUrl, err := url.Parse(parsedLink)
	if err != nil {
		return nil, err
	}

	currHost := currUrl.Host
	if currHost != waveMusicHost {
		return nil, errors.New("invalid host: " + currHost)
	}

	fmt.Println("resultUrl=", parsedLink)
	return currUrl, nil
}

func (useCase linkerUseCase) Create(link string) (string, error) {
	u, err := linkToFormat(link)
	if err != nil {
		return "", err
	}

	returnedHash, err := useCase.linkerAgent.Create(u.String())

	return returnedHash, err
}

func (useCase linkerUseCase) Get(hash string) (string, error) {
	link, err := useCase.linkerAgent.Get(hash)
	if err != nil {
		return "", err
	}

	u, err := url.ParseRequestURI(link)
	if err != nil {
		return "", err
	}

	if !govalidator.IsURL(link) {
		return "", errors.New("invalid link: " + link)
	}

	return u.String(), err
}

func (useCase linkerUseCase) Count(hash string) (int64, error) {
	count, err := useCase.linkerAgent.Count(hash)
	if err != nil {
		return -1, err
	}

	return count, err
}
