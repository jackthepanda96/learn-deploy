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
	//add crazy fitur 
	//yang inggak ada di rancangan
	//sukses membuat anggota project kewalahan
	//wkwk
}
func (ur *userRepository) GetByID(idUser int) (domain.User, error) {
	//hello darkness my old friend

}
