package users

import (
	models "blockbustermvc/internal/models/user"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo models.IUserRepository
}

func NewUserService(userRepo models.IUserRepository) models.IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u UserService) CreateUser(user *models.CreateUserDTO) error {
	if err := user.Validate(); err != nil {
		return nil
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.userRepo.CreateUser(user)
}

func (u UserService) GetUser(id uuid.UUID) (*models.UserDTO, error) {
	return u.userRepo.GetUser(id)
}

func (u UserService) GetAllUsers() ([]*models.UserDTO, error) {
	return u.userRepo.GetAllUsers()
}

func (u UserService) UpdateUser(id uuid.UUID, user *models.UserDTO) error {
	user.UpdatedAt = time.Now()

	return u.userRepo.UpdateUser(id, user)
}

func (u UserService) DeleteUser(id uuid.UUID) error {
	return u.userRepo.DeleteUser(id)
}
