package repositories

import (
	"service-golang/models"

	"gorm.io/gorm"
)

type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error{
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAllUser()([]models.User, error){
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil{
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id uint)(*models.User, error){
	var user models.User
	if err := r.DB.Where("id=?", id).First(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User, id uint) error{
	return r.DB.Model(&models.User{}).Where("id=?", id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error{
	return r.DB.Delete(&models.User{}, id).Error
}