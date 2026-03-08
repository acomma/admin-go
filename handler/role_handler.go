package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/acomma/admin/service"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(database *sql.DB) *RoleHandler {
	return &RoleHandler{
		roleService: service.NewRoleService(database),
	}
}

func (roleHandler *RoleHandler) GetRoleByRoleId(w http.ResponseWriter, r *http.Request) {
	roleId, err := strconv.ParseInt(r.PathValue("roleId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role, err := roleHandler.roleService.GetRoleByRoleId(roleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "角色不存在", http.StatusNotFound)
			return
		}
		http.Error(w, "获取角色失败"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
