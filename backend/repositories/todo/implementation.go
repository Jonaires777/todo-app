package todo

import (
	"errors"
	"time"
	"todo-app/models"

	"gorm.io/gorm"
)

func NewTodoRepository(conn *gorm.DB) Repository {
	return Repository{
		GetBD: func() *gorm.DB {
			return conn
		},
	}
}

func (r *Repository) Create(createData ICreate) (*models.Todos, error) {
	var todo = models.Todos{
		Title:     createData.Title,
		Status:    createData.Status,
		Deadline:  createData.Deadline,
		CreatedBy: createData.CreatedBy,
	}

	result := r.GetBD().Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (r *Repository) ReadAll() ([]models.Todos, error) {
	var todoArray []models.Todos

	result := r.GetBD().Find(&todoArray)

	if result.Error != nil {
		return nil, result.Error
	}

	return todoArray, nil
}

func (r *Repository) Update(updateData IUpdate) (*models.Todos, error) {
	var todo = models.Todos{ID: updateData.ID}
	var fieldMap = make(map[string]interface{})

	if updateData.Title == nil &&
		updateData.Status == nil &&
		updateData.CreatedBy == nil &&
		updateData.Deadline == nil {
		return nil, errors.New("No fields to update")
	}

	if updateData.Title != nil{
		fieldMap["title"] = *updateData.Title
	}

	if updateData.Status != nil{
		fieldMap["status"] = *updateData.Status
	}

	if updateData.CreatedBy != nil{
		fieldMap["createdby"] = *updateData.CreatedBy
	}

	if updateData.Deadline != nil{
		fieldMap["deadline"] = *updateData.Deadline
	}

	fieldMap["updated_at"] = time.Now()

	result := r.GetBD().Model(&todo).Updates(fieldMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (r *Repository) Delete(deleteData IDelete) (*models.Todos, error) {
	var todo = models.Todos{ID: deleteData.ID}

	verifyExistence := r.GetBD().First(&todo)

	if verifyExistence.Error != nil {
		return nil, verifyExistence.Error
	}

	result := r.GetBD().Delete(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}