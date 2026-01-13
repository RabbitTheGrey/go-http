package repository

import (
	"database/sql"
	"errors"
	"go-web/internal/entity"
	"go-web/pkg/db"
	"strings"
)

type UserRepository struct {
	db *sql.DB
	db.DataMapper
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.GetDB()}
}

// Вставка новых данных
func (repo *UserRepository) Insert(entities ...entity.User) error {
	var params []any
	var rows []string

	if len(entities) == 0 {
		return errors.New("Empty data to insert")
	}

	sql := "insert into user (login, password) values "

	for _, entity := range entities {
		rows = append(rows, "(?, ?)")
		params = append(params, entity.Login, entity.Password)
	}

	sql += strings.Join(rows, ", ")

	_, err := repo.db.Exec(sql, params...)

	return err
}

// Обновление пользователя.
// Обновляемые поля: token
func (repo *UserRepository) Update(entity entity.User) error {
	params := []any{entity.Token, entity.ID}

	sql := "update user set token = ? where id = ?"

	_, err := repo.db.Exec(sql, params...)

	return err
}

// Поиск по логину и паролю.
// Может вернуть ошибку ErrNoRows
func (repo *UserRepository) FindByLogin(login string) (entity.User, error) {
	var entity entity.User

	sql := "select * from user where login = ?"

	row := repo.db.QueryRow(sql, login)
	err := repo.ScanRow(&entity, row)

	return entity, err
}

// Поиск по токену
func (repo *UserRepository) FindByToken(token string) (entity.User, error) {
	var entity entity.User

	sql := "select * from user where token = ?"

	row := repo.db.QueryRow(sql, token)
	err := repo.ScanRow(&entity, row)

	return entity, err
}
