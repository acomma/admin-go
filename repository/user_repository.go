package repository

import (
	"database/sql"
	"log"

	"github.com/acomma/admin/model"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (userRepository *UserRepository) GetUserByUserId(userId int64) (model.User, error) {
	var user model.User

	row := userRepository.database.QueryRow("SELECT * FROM t_user WHERE id = ?", userId)
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		log.Printf("查询用户信息失败：%v\n", err)
		return user, err
	}

	return user, nil
}
