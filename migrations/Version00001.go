package migrations

import (
	"database/sql"
)

// Create `User` table
func Version00001(tx *sql.Tx) error {
	sql := `create table if not exists user (
		id integer primary key autoincrement,
		login text not null,
		password text not null,
		token text,
		
		unique (login)
	);`

	_, err := tx.Exec(sql)
	return err
}
