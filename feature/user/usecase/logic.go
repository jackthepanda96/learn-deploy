package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackthepanda96/learn-deploy/domain"
)

type userUseCase struct {
	repository domain.UserRepository
	validator  *validator.Validate
}

func New(r domain.UserRepository, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		repository: r,
		validator:  v,
	}
}

func (uc *userUseCase) Register(newUser domain.User) (domain.User, error) {

}
func (uc *userUseCase) Login(email string, password string) (domain.User, error) {

}
func (uc *userUseCase) Profile(idUser int) (domain.User, error) {

}
