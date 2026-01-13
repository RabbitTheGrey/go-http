package migrations

import (
	"database/sql"
)

// Create `Post` table
func Version00002(tx *sql.Tx) error {
	sql := `create table post (
			id integer primary key autoincrement,
			user_id integer not null,
			title text not null,
			description text not null,
			content text not null,
			created_at text default (datetime('now', 'localtime')),
			updated_at text default (datetime('now', 'localtime')),
			deleted_at text default null,

			unique (title),
			foreign key(user_id) references user(id)
		);

		create trigger update_post_updated_at before update on post
		for each row
		begin
			update post set updated_at = datetime('now', 'localtime') where id = old.id;
		end;`

	_, err := tx.Exec(sql)
	return err
}
