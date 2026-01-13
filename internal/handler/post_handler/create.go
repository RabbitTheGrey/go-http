package post_handler

import (
	"encoding/json"
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/service"
	"go-web/pkg/security"
	"net/http"
)

// Создание поста
//
// Принимает запрос (body):
//   - PostRequest{}
//
// Возвращает:
//   - PostDTO
func Create(w http.ResponseWriter, r *http.Request) {
	var body dto.PostRequest

	user, err := security.CurrentUser(r)
	if err != nil {
		internal.JsonResponse("Forbidden", http.StatusForbidden, w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		internal.JsonResponse("Invalid data provided.", http.StatusBadRequest, w)
		return
	}

	result, serviceErr := service.NewPostService().Create(body, user)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
