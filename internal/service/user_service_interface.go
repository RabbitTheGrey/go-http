package service

type IUserService interface {
	// Создание пользователя
	CreateUser(login string, password string)
}
