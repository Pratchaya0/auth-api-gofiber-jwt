package repositories

import "github.com/Pratchaya0/auth-api-gofiber-jwt/entities"

type UserRepository interface {
	GetUserById(issuer string) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	List() ([]entities.User, error)
	Save(user entities.User) (entities.User, error)
}
