package service

import (
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/entity"
)

type IPostService interface {
	// Получение списка постов с пагинацией, поиском и сортировкой
	List(page int, count int, title string, sort string, direction string) ([]dto.PostDTO, *internal.Error)

	// Получение поста по id
	Read(id int) (dto.PostDTO, *internal.Error)

	// Добавление поста
	Create(body dto.PostRequest, user entity.User) (dto.PostDTO, *internal.Error)

	// Обновление поста
	Update(id int, body dto.PostRequest, user entity.User) (dto.PostDTO, *internal.Error)

	// Удаление поста
	Delete(id int, user entity.User) (bool, *internal.Error)
}
