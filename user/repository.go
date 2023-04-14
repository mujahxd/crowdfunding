package user

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
		return user, err
	}
	return user, nil
}
