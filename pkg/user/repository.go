package user

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *User) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func (ur *gormUserRepository) AddUser(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	log.Printf("Repo : %s", user.FullName)
	db := ur.db.WithContext(ctx).Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) *gormUserRepository {
	return &gormUserRepository{db: db}
}
