package post_handler

import (
	"go-web/internal"
	"go-web/internal/service"
	"go-web/pkg/security"
	"net/http"
	"strconv"
)

// Удаление поста
//
// Принимает параметры (query):
//   - id int
//
// Возвращает:
//   - bool статус удаления
func Delete(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		internal.JsonResponse("Invalid query param `id`", http.StatusBadRequest, w)
		return
	}

	user, err := security.CurrentUser(r)
	if err != nil {
		internal.JsonResponse("Forbidden", http.StatusForbidden, w)
		return
	}

	result, serviceErr := service.NewPostService().Delete(id, user)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
