package post_handler

import (
	"go-web/internal"
	"go-web/internal/entity"
	"go-web/internal/service"
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
	var user entity.User

	ctxUser := r.Context().Value("user")
	user, ok := ctxUser.(entity.User)

	if !ok {
		http.Error(w, "Invalid user type in context", http.StatusInternalServerError)
		return
	}

	queryParams := r.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		internal.JsonResponse("Invalid query param `id`", http.StatusBadRequest, w)
		return
	}

	result, serviceErr := service.NewPostService().Delete(id, user)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
