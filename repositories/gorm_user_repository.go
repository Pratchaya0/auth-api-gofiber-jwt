package repositories

import (
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetUserById(issuer string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", issuer).First(&user).Error

	return user, err
}

func (r *GormUserRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *GormUserRepository) List() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Raw("SELECT id, created_at, updated_at, deleted_at, name, email, password FROM users").Scan(&users).Error
	return users, err
}

func (r *GormUserRepository) Save(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
