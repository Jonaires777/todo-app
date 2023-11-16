package todo

import (
	"database/sql"
)

type (
	Repository struct{
		GetBD *sql.DB
	}

	IInsert struct {
		Title   string
		Content string
		Status  uint
	}

	IDelete struct {
		ID int
	}

	Interface interface {
		Create(IInsert) (error)
		Delete(IDelete) (error)
	}
)