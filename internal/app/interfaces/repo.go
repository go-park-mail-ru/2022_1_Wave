package utilsInterfaces

type RepoInterface interface {
	Insert(Domain) error
	Update(Domain) error
	Delete(int) error
	SelectByID(int) (Domain, error)
	GetAll() ([]Domain, error)
	GetPopular() ([]Domain, error)
	//GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)
}
