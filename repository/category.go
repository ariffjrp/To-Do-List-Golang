package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var data []entity.Category
	dataTerakhir := r.db.WithContext(ctx).Where("user_id = ?", id).Model(&entity.Category{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return []entity.Category{}, nil
	  }else{
		return []entity.Category{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	dataTerakhir := r.db.WithContext(ctx).Create(&category)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return 0, Alert
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	dataTerakhir := r.db.WithContext(ctx).Create(&categories)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return Alert
   }
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var data entity.Category
	dataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).Model(&entity.Category{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return entity.Category{}, nil
	  }else{
		return entity.Category{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	dataTerakhir := r.db.WithContext(ctx).Model(&category).
	Updates(category)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return  Alert
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	DataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).
		Delete(&entity.Category{})
	Alert := DataTerakhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}
