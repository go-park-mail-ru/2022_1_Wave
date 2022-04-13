package utils

import (
	"errors"
)

func ToString(value interface{}) (string, error) {
	str, ok := value.(string)

	if !ok {
		return "", errors.New("error to casting to string")
	}

	return str, nil
}

func ToBool(value interface{}) (bool, error) {
	isBool, ok := value.(bool)

	if !ok {
		return false, errors.New("error to casting to bool")
	}

	return isBool, nil
}

func ToUint64(value interface{}) (uint64, error) {
	float, ok := value.(float64)

	if !ok {
		return 0, errors.New("error to casting to ToUint64")
	}

	unsigned := uint64(float)

	return unsigned, nil
}

func ToInt(value interface{}) (int, error) {
	float, ok := value.(float64)

	if !ok {
		return 0, errors.New("error to casting to ToInt")
	}

	integer := int(float)

	return integer, nil
}

func ToInt64(value interface{}) (int64, error) {
	float, ok := value.(float64)

	if !ok {
		return 0, errors.New("error to casting to ToInt64")
	}

	integer := int64(float)

	return integer, nil
}
