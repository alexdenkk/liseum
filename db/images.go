package db

import (
	"alexdenkk/liseum/model"
	"context"
)

// GetImagesFor - function for getting images records by class ID
func (r *Repository) GetImagesFor(ctx context.Context, classID uint) ([]model.Image, error) {
	var images []model.Image
	result := r.DB.Where("class_id = ?", classID).Find(&images)
	return images, result.Error
}

// CreateImage - function for creating image record
func (r *Repository) CreateImage(ctx context.Context, img *model.Image) error {
	result := r.DB.Create(img)
	return result.Error
}

// CreateImage - function for updating image record
func (r *Repository) UpdateImage(ctx context.Context, img *model.Image) error {
	result := r.DB.Save(img)
	return result.Error
}

func (r *Repository) DeleteImage(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Image{}, id)
	return result.Error
}
