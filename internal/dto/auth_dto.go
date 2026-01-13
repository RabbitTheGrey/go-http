package dto

import "go-web/internal/entity"

type AuthDTO struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func (dto *AuthDTO) CreateFromEntity(entity entity.User) {
	dto.ID = entity.ID
	dto.Token = *entity.Token
}
