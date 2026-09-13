package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() (*UserRepo, error) {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}

	return &UserRepo{db: db}, nil
}

func (r *UserRepo) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepo) GetByID(id uint64) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepo) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint64) error {
	return r.db.Delete(&User{}, id).Error
}