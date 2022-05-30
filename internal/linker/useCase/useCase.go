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
}

type linkerUseCase struct {
	linkerAgent domain.LinkerAgent
}

func NewLinkerUseCase(linkerAgent domain.LinkerRepo) *linkerUseCase {
	return &linkerUseCase{
		linkerAgent: linkerAgent,
	}
}

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

	wwwPart := "://www."

	wwwBegin := strings.Index(parsedLink, wwwPart)
	if wwwBegin != -1 {
		wwwEnd := wwwBegin + len(wwwPart)
		parsedLink = parsedLink[:wwwBegin+3] + parsedLink[wwwEnd:]
	}

	fmt.Println("resultUrl=", parsedLink)

	return url.Parse(parsedLink)
}

func (useCase linkerUseCase) Create(link string) (string, error) {
	u, err := linkToFormat(link)
	if err != nil {
		return "", err
	}

	returnedHash, err := useCase.linkerAgent.Create(u.String())

	checkedUrl := url.URL{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		User:        u.User,
		Host:        "wave-music.xyz",
		Path:        returnedHash,
		RawPath:     returnedHash,
		ForceQuery:  u.ForceQuery,
		RawQuery:    u.RawQuery,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
	}

	fmt.Println(checkedUrl.String())

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

	fmt.Println("result url=", u.String())
	return u.String(), err
}
