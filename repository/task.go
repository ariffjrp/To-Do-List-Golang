package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var data []entity.Task
	dataTerakhir := r.db.WithContext(ctx).Where("user_id = ?", id).Model(&entity.Task{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	  }else{
		return []entity.Task{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	dataTerakhir := r.db.WithContext(ctx).Create(&task)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return 0, Alert
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var data entity.Task
	dataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).Model(&entity.Task{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return entity.Task{}, nil
	  }else{
		return entity.Task{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var data []entity.Task
	dataTerakhir := r.db.WithContext(ctx).Where("category_id = ?", catId).Model(&entity.Task{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	  }else{
		return []entity.Task{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	dataTerakhir := r.db.WithContext(ctx).Model(&task).
	Where("user_id", task.UserID).
	Updates(task)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	DataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).
		Delete(&entity.Task{})
	Alert := DataTerakhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}
