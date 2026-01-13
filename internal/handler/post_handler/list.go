package post_handler

import (
	"go-web/internal"
	"go-web/internal/service"
	"net/http"
	"strconv"
)

const (
	defaultPage  = 1
	defaultCount = 10
)

// Просмотр списка постов
//
// Принимает параметры (query):
//   - page int (default 1)
//   - count int (default 10)
//   - title string (search)
//   - sort (order by)
//   - direction (order direction)
//
// Возвращает:
//   - []PostDTO
func List(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	// Простой и компактный способ извлечения числового параметра, но вместо ошибки подставляет значение по умолчанию
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		page = defaultPage
	}

	count, err := strconv.Atoi(queryParams.Get("count"))
	if err != nil {
		count = defaultCount
	}

	title := queryParams.Get("title")
	sort := queryParams.Get("sort")
	direction := queryParams.Get("direction")

	result, serviceErr := service.NewPostService().List(page, count, title, sort, direction)
	if serviceErr != nil {
		internal.JsonResponse(serviceErr.Message, serviceErr.Code, w)
		return
	}

	internal.JsonResponse(result, http.StatusOK, w)
}
