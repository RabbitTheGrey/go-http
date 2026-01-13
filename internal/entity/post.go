package entity

type Post struct {
	ID          int     `db:"id"`
	UserID      int     `db:"user_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Content     string  `db:"content"`
	CreatedAt   *string `db:"created_at"` //time.Time для других бд
	UpdatedAt   *string `db:"updated_at"` //time.Time для других бд
	DeletedAt   *string `db:"deleted_at"` //time.Time для других бд
}
