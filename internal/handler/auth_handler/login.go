package auth_handler

import (
	"encoding/json"
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/repository"
	"go-web/pkg/security"
	"net/http"
)

// Аутентификация
//
// Принимает параметры (body):
//   - dto.AuthRequest
//
// Возвращает:
//   - dto.AuthDTO
func Login(w http.ResponseWriter, r *http.Request) {
	var data dto.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		internal.JsonResponse("Invalid data provided.", http.StatusBadRequest, w)
		return
	}

	repository := repository.NewUserRepository()
	currentUser, err := repository.FindByLogin(data.Login)
	if err != nil {
		internal.JsonResponse("Invalid credentials.", http.StatusBadRequest, w)
		return
	}

	err = security.ValidatePassword(currentUser, data.Password)
	if err != nil {
		internal.JsonResponse(err, http.StatusBadRequest, w)
		return
	}

	var isTokenValid bool

	if currentUser.Token != nil {
		isTokenValid, err = security.IsTokenValid(*currentUser.Token)
	} else {
		err = security.RefreshToken(&currentUser)
	}

	if err != nil || !isTokenValid {
		if err != nil {
			internal.JsonResponse(err, http.StatusBadRequest, w)
			return
		}
		repository.Update(currentUser)
	}

	dto := dto.AuthDTO{}
	dto.CreateFromEntity(currentUser)

	internal.JsonResponse(dto, http.StatusOK, w)
}
