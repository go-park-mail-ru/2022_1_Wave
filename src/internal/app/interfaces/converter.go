package utilsInterfaces

import "reflect"

type ConverterInterface interface {
	ToDomains(objects []interface{}) (*[]Domain, error)
	ToDataTransfers(toType reflect.Type, objects []interface{}) (*[]DataTransfer, error)
	ToString(value interface{}) (string, error)
	ToUint64(value interface{}) (uint64, error)
	ToInt(value interface{}) (int, error)
	GetDataTransferTypeByDomainType(domainType reflect.Type) (reflect.Type, error)
}
