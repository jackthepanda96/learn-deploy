package repository

import (
	"github.com/jackthepanda96/learn-deploy/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
func (ur *userRepository) Insert(newUser domain.User) (domain.User, error) {
	// Edit by Atha DF
	return domain.User{}, nil
	// Last edit
}
func (ur *userRepository) GetForLogin(email string, password string) (domain.User, error) {
	//line 1
	//line 2
	//line 99
	//dari jerry
	//nambah lagi
}
func (ur *userRepository) GetByID(idUser int) (domain.User, error) {

}
