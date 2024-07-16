package level

import "gorm.io/gorm"

type Repository interface {
	InputLevel(level Level)(Level, error)
	FindAll()([]Level, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InputLevel(level Level)(Level, error) {
	err := r.db.Create(&level).Error
	if err != nil {
		return level, err
	}
	return level, nil
}

func (r *repository) FindAll()([]Level, error) {
	var level []Level
	err := r.db.Find(&level).Error
	if err != nil {
		return nil, err
	}
	return level, nil
}