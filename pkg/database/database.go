package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

func InitDb(dsn string) (*Database, error) {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &Database{
		conn: conn,
	}, err
}
