package linkerDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	LinkerUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/linker/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/linkerProto"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	LinkerUseCase LinkerUseCase.LinkerUseCase
}

func MakeHandler(linker LinkerUseCase.LinkerUseCase) Handler {
	return Handler{
		LinkerUseCase: linker,
	}
}

// Get godoc
// @Summary      Get
// @Description  getting hash by url
// @Tags         linker
// @Produce      application/json
// @Param        hash   path      string  true  "hash to be converted to url"
// @Success      301  {object}  webUtils.Success
// @Failure      404  {object}  webUtils.Error  "Not found"
// @Router       /api/v1/linker/{hash} [get]
func (h Handler) Get(ctx echo.Context) error {
	hash := ctx.Param(internal.Hash)
	url, err := h.LinkerUseCase.Get(hash)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusNotFound)
	}

	return ctx.JSON(http.StatusMovedPermanently,
		webUtils.Success{
			Status: webUtils.OK,
			Result: url})
}

// Create godoc
// @Summary      Create
// @Description  creating hash by url or return existing hash
// @Tags         linker
// @Accept          application/json
// @Produce      application/json
// @Param        Url  body      linkerProto.UrlWrapper  true  "url to hash"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Router       /api/v1/linker/ [post]
func (h Handler) Create(ctx echo.Context) error {
	url := linkerProto.UrlWrapper{}

	if err := ctx.Bind(&url); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	hash, err := h.LinkerUseCase.Create(url.Url)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: linkerProto.HashWrapper{Hash: hash}})
}
