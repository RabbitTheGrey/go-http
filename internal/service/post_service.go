package service

import (
	"go-web/internal"
	"go-web/internal/dto"
	"go-web/internal/entity"
	"go-web/internal/repository"
)

type PostService struct {
	repository repository.PostRepository
}

func NewPostService() IPostService {
	return &PostService{
		repository: *repository.NewPostRepository(),
	}
}

func (s *PostService) List(page int, count int, title string, sort string, direction string) ([]dto.PostDTO, *internal.Error) {
	var listDTO []dto.PostDTO
	list, err := s.repository.FindAll(page, count, title, sort, direction)
	if err != nil {
		return listDTO, internal.ErrInternal(err.Error())
	}

	for _, post := range list {
		dto := dto.PostDTO{}
		dto.CreateFromEntity(post)
		listDTO = append(listDTO, dto)
	}

	return listDTO, nil
}

func (s *PostService) Read(id int) (dto.PostDTO, *internal.Error) {
	post, err := s.repository.Find(id)
	if err != nil {
		return dto.PostDTO{}, internal.ErrNotFound("Пост не найден")
	}

	dto := dto.PostDTO{}
	dto.CreateFromEntity(post)

	return dto, nil
}

func (s *PostService) Create(body dto.PostRequest, user entity.User) (dto.PostDTO, *internal.Error) {
	post := entity.Post{
		UserID:      user.ID,
		Title:       body.Title,
		Description: body.Description,
		Content:     body.Content,
	}

	err := s.repository.Insert(post)
	if err != nil {
		return dto.PostDTO{}, internal.ErrInternal(err.Error())
	}

	dto := dto.PostDTO{}
	dto.CreateFromEntity(post)

	return dto, nil
}

func (s *PostService) Update(id int, body dto.PostRequest, user entity.User) (dto.PostDTO, *internal.Error) {
	post, err := s.repository.Find(id)
	if err != nil {
		return dto.PostDTO{}, internal.ErrNotFound("Пост не найден")
	}

	if post.UserID != user.ID {
		return dto.PostDTO{}, internal.ErrForbidden("Вы не являетесь автором поста.")
	}

	post.Title = body.Title
	post.Description = body.Description
	post.Content = body.Content

	err = s.repository.Update(post)
	if err != nil {
		return dto.PostDTO{}, internal.ErrInternal(err.Error())
	}

	dto := dto.PostDTO{}
	dto.CreateFromEntity(post)

	return dto, nil
}

func (s *PostService) Delete(id int, user entity.User) (bool, *internal.Error) {
	post, err := s.repository.Find(id)
	if err != nil {
		return false, internal.ErrNotFound("Пост не найден")
	}

	if post.UserID != user.ID {
		return false, internal.ErrForbidden("Вы не являетесь автором поста.")
	}

	err = s.repository.Delete(post)
	if err != nil {
		return false, internal.ErrInternal(err.Error())
	}

	return true, nil
}
