package post_handler

import (
	"encoding/json"
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/service"
	"go-web/pkg/security"
	"net/http"
	"strconv"
)

// Редактирование поста
//
// Принимает параметры (query):
//   - id int
//
// Принимает запрос (body):
//   - PostRequest{}
//
// Возвращает:
//   - PostDTO
func Update(w http.ResponseWriter, r *http.Request) {
	var body dto.PostRequest

	user, err := security.CurrentUser(r)
	if err != nil {
		internal.JsonResponse("Forbidden", http.StatusForbidden, w)
		return
	}

	queryParams := r.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		internal.JsonResponse("Invalid query param `id`", http.StatusBadRequest, w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		internal.JsonResponse("Invalid data provided.", http.StatusBadRequest, w)
		return
	}

	result, serviceErr := service.NewPostService().Update(id, body, user)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
