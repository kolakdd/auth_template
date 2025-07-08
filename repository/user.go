// Package repository
package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kolakdd/auth_template/models"
	"github.com/kolakdd/auth_template/query"
	"gorm.io/gorm"
)

type RepositoryUser interface {
	NewUser(name string) (*models.User, error)
	GetUser(id uuid.UUID) (*models.User, error)
	DeactivateUser(id uuid.UUID) (*models.User, error)
}
type repositoryUser struct {
	db *gorm.DB
}

func NewRepoUser(collection *gorm.DB) RepositoryUser {
	return &repositoryUser{
		db: collection,
	}
}

func (r *repositoryUser) NewUser(name string) (*models.User, error) {
	q := query.Use(r.db)
	userDB := models.NewUserDB(name)
	if err := q.User.Create(&userDB); err != nil {
		return nil, err
	}
	return &userDB, nil
}

func (r *repositoryUser) GetUser(id uuid.UUID) (*models.User, error) {
	user := new(models.User)
	err := r.db.Where("GUID = ?", id).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found ")
	}
	if user.Deactivated {
		return nil, fmt.Errorf("user deactivated ")
	}
	return user, nil
}

func (r *repositoryUser) DeactivateUser(id uuid.UUID) (*models.User, error) {
	user := new(models.User)
	err := r.db.Model(&user).Where("guid = ?", id).Update("deactivated", true).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
