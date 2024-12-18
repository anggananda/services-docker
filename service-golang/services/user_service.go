package services

import (
	"service-golang/models"
	"service-golang/repositories"
)

type UserService struct{
	UserRepository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService{
	return &UserService{UserRepository: repo}
}

func (s *UserService) CreateUser(user *models.User) error{
	return s.UserRepository.CreateUser(user)
}

func (s *UserService) GetAllUser()([]models.User, error){
	return s.UserRepository.GetAllUser()
}

func (s *UserService) GetUserByID(id uint)(*models.User, error){
	return s.UserRepository.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User, id uint) error{
	return s.UserRepository.UpdateUser(user, id)
}

func (s *UserService) DeleteUser(id uint) error{
	return s.UserRepository.DeleteUser(id)
}