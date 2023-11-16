package todo

import "database/sql"

func NewTodoRepository(db *sql.DB) Repository {
	return Repository{
		GetBD: db,
	}
}