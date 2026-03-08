package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/acomma/admin/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(database *sql.DB) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(database),
	}
}

func (userHandler *UserHandler) GetUserByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userHandler.userService.GetUserByUserId(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "用户不存在", http.StatusNotFound)
			return
		}
		http.Error(w, "获取用户失败"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
