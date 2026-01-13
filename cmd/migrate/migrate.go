package main

import (
	"database/sql"
	"fmt"
	"go-web/conf"
	"go-web/migrations"
	"go-web/pkg/helper"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

type migration struct {
	version string
	execute func(tx *sql.Tx) error
}

// Список миграций. Добавлять новые миграции сюда
func buildMigrationsList() []migration {
	var migrationsList []migration

	migrationsList = append(migrationsList,
		migration{"Version00001", migrations.Version00001},
		migration{"Version00002", migrations.Version00002},
		//...
	)

	return migrationsList
}

// Проверяет, применена ли данная версия миграции
func isMigrationExists(version string) bool {
	var isExists uint8
	sql := `select count(*) from migrations where version = ?`

	err := database.QueryRow(sql, version).Scan(&isExists)
	helper.HandlePanic(err)

	return isExists != 0
}

// Сохранение примененной версии в таблице migrations
func commitVersion(verion string, tx *sql.Tx) error {
	sql := `insert into migrations (version) values (?);`
	_, err := tx.Exec(sql, verion)

	return err
}

// Проверка и выполнение миграции
func checkAndExecute(migration migration) {
	if isMigrationExists(migration.version) {
		fmt.Printf("Migration version '%s' already exists. Skipping...\n", migration.version)
		return
	}

	tx, err := database.Begin()
	helper.HandlePanic(err)

	defer func() {
		if tx != nil {
			tx.Rollback()
		}
	}()

	err = migration.execute(tx)
	helper.HandlePanic(err)

	err = commitVersion(migration.version, tx)
	helper.HandlePanic(err)

	tx.Commit()
	tx = nil
	fmt.Printf("Migration version '%s' successfully applied!\n", migration.version)
}

// Создание таблицы миграций
func createMigrationsTable() {
	sql := `create table if not exists migrations(
		version text primary key,
		executed_at text default (datetime('now', 'localtime'))
	);`

	_, err := database.Exec(sql)
	helper.HandlePanic(err)
}

func main() {
	migrationsList := buildMigrationsList()
	var err error

	database, err = sql.Open(conf.DbDriver, conf.DbConnectionString)
	helper.HandlePanic(err)
	defer database.Close()

	createMigrationsTable()

	for _, migration := range migrationsList {
		checkAndExecute(migration)
	}
}
