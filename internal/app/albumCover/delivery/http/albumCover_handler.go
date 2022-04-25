package albumCoverDeliveryHttp

//type Handler struct {
//	albumCoverUseCase AlbumCoverUseCase.AlbumCoverUseCase
//}
//
//func MakeHandler(cover AlbumCoverUseCase.AlbumCoverUseCase) Handler {
//	return Handler{
//		albumCoverUseCase: cover,
//	}
//}

// GetAll godoc
// @Summary      GetAll
// @Description  getting all albums cover
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [get]
//func (h Handler) GetAll(ctx echo.Context) error {
//	domains, err := h.albumCoverUseCase.GetAll()
//
//	if err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	if domains == nil {
//		domains = []domain.AlbumCoverDataTransfer{}
//	}
//
//	return ctx.JSON(http.StatusOK,
//		webUtils.Success{
//			Status: webUtils.OK,
//			Result: domains})
//}

// Create godoc
// @Summary      Create
// @Description  creating new albumCover
// @Tags         albumCover
// @Accept       application/json
// @Produce      application/json
// @Param        Album  body      domain.AlbumCover  true  "params of new album cover. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [post]
//func (h Handler) Create(ctx echo.Context) error {
//	result := domain.AlbumCover{}
//
//	if err := ctx.Bind(&result); err != nil {
//		return err
//	}
//
//	if err := result.Check(); err != nil {
//		return err
//	}
//
//	if err := h.albumCoverUseCase.Create(result); err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	lastId, err := h.albumCoverUseCase.GetLastId()
//	if err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	return ctx.JSON(http.StatusOK,
//		webUtils.Success{
//			Status: webUtils.OK,
//			Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
//}

// Update godoc
// @Summary      Update
// @Description  updating album cover by id
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Param        AlbumCover  body      domain.AlbumCover  true  "id of updating album cover and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [put]
//func (h Handler) Update(ctx echo.Context) error {
//	result := domain.AlbumCover{}
//
//	if err := ctx.Bind(&result); err != nil {
//		return err
//	}
//
//	if err := result.Check(); err != nil {
//		return err
//	}
//
//	if err := h.albumCoverUseCase.Update(result); err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	id := result.Id
//	return ctx.JSON(http.StatusOK,
//		webUtils.Success{
//			Status: webUtils.OK,
//			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
//}

// Get godoc
// @Summary      Get
// @Description  getting album cover by id
// @Tags         albumCover
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album cover which need to be getted"
// @Success      200  {object}  domain.AlbumCover
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/{id} [get]
//func (h Handler) Get(ctx echo.Context) error {
//	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
//	if err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//	if id < 0 {
//		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
//	}
//
//	album, err := h.albumCoverUseCase.GetById(int64(id))
//
//	if err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	return ctx.JSON(http.StatusOK,
//		webUtils.Success{
//			Status: webUtils.OK,
//			Result: album})
//}

// Delete godoc
// @Summary      Delete
// @Description  deleting album cover by id
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album cover which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/{id} [delete]
//func (h Handler) Delete(ctx echo.Context) error {
//	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
//	if err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//	if id < 0 {
//		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
//	}
//
//	if err := h.albumCoverUseCase.Delete(int64(id)); err != nil {
//		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
//	}
//
//	return ctx.JSON(http.StatusOK,
//		webUtils.Success{
//			Status: webUtils.OK,
//			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
//}
