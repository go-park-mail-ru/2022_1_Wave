package utils

import (
	"errors"
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
