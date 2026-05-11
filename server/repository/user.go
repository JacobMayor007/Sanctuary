package repository

import (
	"sanctuary_server/database"
	"sanctuary_server/model"
)

// Interface
type UserRepositoryInterface interface {
	CreateUser(firebaseUID, email, name string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(firebaseUID, name string) (*model.User, error)
	DeleteUser(firebaseUID string) error
}

// Concrete struct
type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(firebaseUID, email, name string) (*model.User, error) {
	var existing model.User
	result := database.DB.Where("f_uid = ?", firebaseUID).First(&existing)
	if result.Error == nil {
		return &existing, nil
	}

	user := model.User{
		FUID:  firebaseUID,
		Email: email,
		Name:  name,
	}

	result = database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(firebaseUID, name string) (*model.User, error) {
	var user model.User
	result := database.DB.Where("f_uid = ?", firebaseUID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	user.Name = name
	database.DB.Save(&user)

	return &user, nil
}

func (u *UserRepository) DeleteUser(firebaseUID string) error {
	result := database.DB.Where("f_uid = ?", firebaseUID).Delete(&model.User{})
	return result.Error
}
