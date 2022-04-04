package utilsInterfaces

type Domain interface {
	GetId() uint64
	SetId(id uint64) (Domain, error)
	Check() error
	GetCountListening() uint64
	CastDomainToDataTransferObject(domain Domain) (DataTransfer, error)
}
