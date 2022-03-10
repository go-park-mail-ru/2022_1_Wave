package db

import "fmt"

func convertToString(info interface{}) string {
	result := ""

	str, ok := info.(string)
	if !ok {
		result = fmt.Sprint(info)
	} else {
		result = str
	}
	return result
}

func SuccessWrapper(info interface{}, wrap string) string {
	return wrap + "(" + convertToString(info) + ")"
}
