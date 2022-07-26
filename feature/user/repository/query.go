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

}
func (ur *userRepository) GetForLogin(email string, password string) (domain.User, error) {
	//line 1
	//line 2
	//line 99
	//dari jerry
}
func (ur *userRepository) GetByID(idUser int) (domain.User, error) {
	// sendok
	// garpu
	// 3 hari lagi hari jumat :(
}
