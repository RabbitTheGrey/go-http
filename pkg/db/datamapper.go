package db

import (
	"database/sql"
	"reflect"
)

type DataMapper struct{}

// Маппинг многострочного SQL-вывода в срез структур
// Принимает указатель на срез структур: &list
func (dm *DataMapper) ScanRows(list any, rows *sql.Rows) error {
	listVal := reflect.ValueOf(list).Elem()
	structType := listVal.Type().Elem()
	fieldsCount := structType.NumField()

	var fieldIndices []int
	for i := range fieldsCount {
		field := structType.Field(i)
		if field.PkgPath != "" {
			continue
		}
		if tag := field.Tag.Get("db"); tag == "" || tag == "-" {
			continue
		}
		fieldIndices = append(fieldIndices, i)
	}

	for rows.Next() {
		structVal := reflect.New(structType).Elem()
		args := make([]any, len(fieldIndices))

		for i, idx := range fieldIndices {
			args[i] = structVal.Field(idx).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}
		listVal.Set(reflect.Append(listVal, structVal))
	}

	return rows.Err()
}

// Маппинг однострочного SQL-вывода в структуру
// Принимает указатель на структуру: &entity
func (dm *DataMapper) ScanRow(entity any, row *sql.Row) error {
	reflectValue := reflect.ValueOf(entity).Elem()
	reflectType := reflectValue.Type()
	args := make([]any, 0, reflectType.NumField())

	for i := range reflectType.NumField() {
		field := reflectValue.Field(i)
		if !field.CanSet() {
			continue
		}
		tag := reflectType.Field(i).Tag.Get("db")
		if tag == "" || tag == "-" {
			continue
		}
		args = append(args, field.Addr().Interface())
	}

	return row.Scan(args...)
}
