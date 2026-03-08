package service

import (
	"database/sql"

	"github.com/acomma/admin/model"
	"github.com/acomma/admin/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(database *sql.DB) *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(database),
	}
}

func (userService *UserService) GetUserByUserId(userId int64) (model.User, error) {
	user, err := userService.userRepository.GetUserByUserId(userId)
	if err != nil {
		return user, err
	}
	return user, nil
}
