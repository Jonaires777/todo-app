package user

import (
	"database/sql"
)

const (
	emailMaxLen    = 128
	emailMinLen    = 3
	nameMaxLen     = 128
	nameMinLen     = 3
	passwordMinLen = 8
)

type (
	Repository struct{
		GetBD *sql.DB
	}

	IInsert struct {
		Name     string
		Email    string
		Password string
	}

	IDelete struct {
		ID int
	}

	Interface interface {
		Insert(IInsert) (error)
		Delete(IDelete) (error)
	}
)