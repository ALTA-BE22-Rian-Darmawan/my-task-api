package user

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(input Core) error
	Delete(id uint) error
	Update(id uint, input Core) error
	SelectByEmail(email string) (*Core, error)
	SelectById(id uint) (*Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	Delete(id uint) error
	Update(id uint, input Core) error
	Login(email, password string) (data *Core, token string, err error)
	GetById(id uint) (data *Core, err error)
}
