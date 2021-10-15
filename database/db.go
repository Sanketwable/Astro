package database

import (
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sql6444259:ZekMxfwtYS@tcp(sql6.freesqldatabase.com:3306)/sql6444259")
	if err != nil {
		return nil, err
	}
	return db, nil

}
