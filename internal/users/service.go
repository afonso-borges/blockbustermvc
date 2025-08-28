package users

import (
	models "blockbustermvc/internal/models/user"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository models.IUserRepository
}

func NewUserService(userRepo models.IUserRepository) models.IUserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (u UserService) CreateUser(user *models.CreateUserDTO) error {
	return u.userRepository.CreateUser(user)
}

func (u UserService) GetUser(id uuid.UUID) (*models.UserDTO, error) {
	return u.userRepository.GetUserById(id)
}

func (u UserService) GetAllUsers() ([]*models.UserDTO, error) {
	return u.userRepository.GetAllUsers()
}

func (u UserService) UpdateUser(id uuid.UUID, user *models.UpdateUserDTO) error {
	return u.userRepository.UpdateUser(id, user)
}

func (u UserService) DeleteUser(id uuid.UUID) error {
	return u.userRepository.DeleteUser(id)
}
