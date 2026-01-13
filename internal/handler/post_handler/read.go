package post_handler

import (
	"go-web/internal"
	"go-web/internal/service"
	"net/http"
	"strconv"
)

// Детальный просмотр поста
//
// Принимает параметры (query):
//   - id int
//
// Возвращает:
//   - PostDTO
func Read(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		internal.JsonResponse("Invalid query param `id`", http.StatusBadRequest, w)
		return
	}

	result, serviceErr := service.NewPostService().Read(id)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
