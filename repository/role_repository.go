package repository

import (
	"database/sql"
	"log"

	"github.com/acomma/admin/model"
)

type RoleRepository struct {
	database *sql.DB
}

func NewRoleRepository(database *sql.DB) *RoleRepository {
	return &RoleRepository{
		database: database,
	}
}

func (roleRepository *RoleRepository) GetRoleByRoleId(roleId int64) (model.Role, error) {
	var role model.Role

	row := roleRepository.database.QueryRow("SELECT * FROM t_role WHERE id = ?", roleId)
	if err := row.Scan(&role.Id, &role.Name); err != nil {
		log.Printf("查询角色信息失败：%v\n", err)
		return role, err
	}

	return role, nil
}
