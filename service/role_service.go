package service

import (
	"database/sql"

	"github.com/acomma/admin/model"
	"github.com/acomma/admin/repository"
)

type RoleService struct {
	roleRepository *repository.RoleRepository
}

func NewRoleService(database *sql.DB) *RoleService {
	return &RoleService{
		roleRepository: repository.NewRoleRepository(database),
	}
}

func (roleService *RoleService) GetRoleByRoleId(roleId int64) (model.Role, error) {
	role, err := roleService.roleRepository.GetRoleByRoleId(roleId)
	if err != nil {
		return role, err
	}
	return role, nil
}
