package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-web/internal/entity"
	"go-web/pkg/db"
	"strings"
)

type PostRepository struct {
	db *sql.DB
	db.DataMapper
}

func NewPostRepository() *PostRepository {
	return &PostRepository{db: db.GetDB()}
}

// Вставка новых данных
func (repo *PostRepository) Insert(entities ...entity.Post) error {
	var params []any
	var rows []string

	if len(entities) == 0 {
		return errors.New("Empty data to insert")
	}

	sql := "insert into post (user_id, title, description, content) values "

	for _, entity := range entities {
		rows = append(rows, "(?, ?, ?, ?)")
		params = append(params,
			entity.UserID,
			entity.Title,
			entity.Description,
			entity.Content,
		)
	}

	sql += strings.Join(rows, ", ")

	_, err := repo.db.Exec(sql, params...)

	return err
}

// Обновление поста
// Обновляемые поля: title, description, content
func (repo *PostRepository) Update(entity entity.Post) error {
	params := []any{
		entity.Title,
		entity.Description,
		entity.Content,
		entity.ID,
	}

	sql := `update post
		set
			title = ?,
			description = ?,
			content = ?
		where
			id = ?`

	_, err := repo.db.Exec(sql, params...)

	return err
}

// Софт-удаление поста
// Метка удаления: deleted_at
func (repo *PostRepository) Delete(entity entity.Post) error {
	sql := "update post set deleted_at = datetime('now', 'localtime') where id = ?"

	_, err := repo.db.Exec(sql, entity.ID)

	return err
}

// Поиск поста по id
func (repo *PostRepository) Find(id int) (entity.Post, error) {
	var entity entity.Post

	sql := "select * from post where id = ? and deleted_at is null"

	row := repo.db.QueryRow(sql, id)
	err := repo.ScanRow(&entity, row)

	return entity, err
}

// Поиск постов по title с пагинацией и сортировкой
func (repo *PostRepository) FindAll(
	page int, // Номер страницы
	count int, // Количество записей на странице
	title string, // Заголовок (поле для поиска)
	sort string, // Поле для сортировки
	direction string, // Направление сортировки ("asc", "desc")
) ([]entity.Post, error) {
	var list []entity.Post
	var params []any

	sql := "select * from post where deleted_at is null"

	if title != "" {
		sql += " and title like ?"
		params = append(params, "%"+title+"%")
	}

	if sort != "" {
		if direction == "" {
			direction = "asc"
		}

		sql += fmt.Sprintf(" order by %s %s", sort, direction)
	}

	sql += fmt.Sprintf(" limit %d offset %d", count, count*(page-1))

	rows, err := repo.db.Query(sql, params...)

	if err != nil {
		return list, err
	}
	defer rows.Close()

	err = repo.ScanRows(&list, rows)

	return list, err
}
