package post_handler

import (
	"encoding/json"
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/entity"
	"go-web/internal/service"
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
	var user entity.User

	ctxUser := r.Context().Value("user")
	user, ok := ctxUser.(entity.User)

	if !ok {
		http.Error(w, "Invalid user type in context", http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&body)
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
