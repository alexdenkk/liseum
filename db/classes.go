package db

import (
	"alexdenkk/liseum/model"
	"context"
)

// GetAllClasses - function for getting all classes records
func (r *Repository) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	var classes []model.Class
	result := r.DB.Find(&classes)
	return classes, result.Error
}

// CreateClass - function for creating Class record
func (r *Repository) CreateClass(ctx context.Context, class *model.Class) error {
	return r.DB.Create(class).Error
}

// UpdateClass - function for updating Class record
func (r *Repository) UpdateClass(ctx context.Context, class model.Class) error {
	return r.DB.Save(&class).Error
}

// DeleteClass - function for deleting class record by id
func (r *Repository) DeleteClass(ctx context.Context, id uint) error {
	return r.DB.Delete(&model.Class{}, id).Error
}

// GetClassByName - function for getting class by name
func (r *Repository) GetClassByName(ctx context.Context, name string) (model.Class, error) {
	var class model.Class
	result := r.DB.Where("name = ?", name).First(&class)
	return class, result.Error
}
