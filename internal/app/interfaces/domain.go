package utilsInterfaces

type Domain interface {
	GetId() int
	SetId(id int) (Domain, error)
	Check() error
	GetCountListening() int
}
