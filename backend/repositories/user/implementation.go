package user

import (
	"database/sql"
	"errors"
	"strings"
	"todo-app/models"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository(db *sql.DB) Repository {
	return Repository{
		GetBD: db,
	}
}

func (r *Repository) Insert(insertData IInsert) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(insertData.Password), 12)
	if err != nil {
		return err
	}
	// Criando o usuário
	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`
	_, err = r.GetBD.Exec(stmt, insertData.Name, insertData.Email, string(hashedPassword))
	
	//Validando 
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}

		return err
	}

	return nil
}

func (r *Repository) Exists(id int) (bool, error) {
	// Checar se o Usuário existe ou não
	
	return false, nil
}

func (r *Repository) Delete(deleteUser IDelete) (error){
	
	// verificando se o usuário existe antes de deletar
	verifyExistence, err := r.Exists(deleteUser.ID)
	if err != nil {
		return err
	}

	if !verifyExistence{
		return models.ErrNoRecord
	}
	
	// Deletar o usuário azendo alguma 
	//	stmt := `DELETE FROM users bla bla bla`
	//	err = r.GetBD.Exec(stmt, deleteUser)

	return nil
}