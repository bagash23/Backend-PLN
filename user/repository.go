package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByUsername(username string)(User, error)
	FindByID(id string)(User, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepostiory(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUsername(username string)(User, error) {
	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(id string)(User, error) {
	var user User	
	err := r.db.Where("id_user = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}