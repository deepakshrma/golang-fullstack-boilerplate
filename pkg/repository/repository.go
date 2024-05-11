package repository

import (
	"database/sql"
	"webapp/pkg/model"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllUsers() ([]*model.User, error)
}
