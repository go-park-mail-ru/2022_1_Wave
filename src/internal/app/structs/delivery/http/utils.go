package structsDeliveryHttp

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
	"github.com/labstack/echo/v4"
	"reflect"
	"strconv"
)

func readPostPutRequest(ctx echo.Context, domainType reflect.Type) (utilsInterfaces.Domain, error) {
	var result interface{}

	if err := ctx.Bind(&result); err != nil {
		return nil, err
	}

	concreteDomain, errDueCast := domainCreator.CreateDomainFromInterface(domainType, result)

	if errDueCast != nil {
		return nil, errDueCast
	}

	if err := concreteDomain.Check(); err != nil {
		return nil, err
	}

	object := concreteDomain

	return object, nil
}

func readGetDeleteRequest(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return constants.BadId, err
	}

	if id < 0 {
		return constants.BadId, errors.New(constants.IndexOutOfRange)
	}

	return id, nil
}
