package dto

import (
	"go-web/internal/entity"
)

type PostDTO struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`    //*time.Time для других бд
	CreatedAt   *string `json:"created_at"` //*time.Time для других бд
	UpdatedAt   *string `json:"updated_at"` //*time.Time для других бд
}

func (dto *PostDTO) CreateFromEntity(entity entity.Post) {
	dto.ID = entity.ID
	dto.UserID = entity.UserID
	dto.Title = entity.Title
	dto.Description = entity.Description
	dto.Content = entity.Content
	dto.CreatedAt = entity.CreatedAt
	dto.UpdatedAt = entity.UpdatedAt
}
