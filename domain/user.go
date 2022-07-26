package domain

type User struct {
	ID       int
	Nama     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserUseCase interface {
	Register(newUser User) (User, error)
	Login(email string, password string) (User, error)
	Profile(idUser int) (User, error)
}

type UserRepository interface {
	Insert(newUser User) (User, error)
	GetForLogin(email string, password string) (User, error)
	GetByID(idUser int) (User, error)
}
