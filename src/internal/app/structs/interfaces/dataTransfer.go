package utilsInterfaces

type DataTransfer interface {
	CreateDataTransferFromInterface(data interface{}) (DataTransfer, error)
}
