package validators

import (
	"testing"

	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	usercase "github.com/Pratchaya0/auth-api-gofiber-jwt/usecases"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	getUserByIdFunc    func(id string) (entities.User, error)
	getUserByEmailFunc func(email string) (entities.User, error)
	listFunc           func() ([]entities.User, error)
	saveFunc           func(user entities.User) (entities.User, error)
}

// GetUserByEmail implements repositories.UserRepository.
func (m *mockUserRepo) GetUserByEmail(email string) (entities.User, error) {
	return m.getUserByEmailFunc(email)
}

// GetUserById implements repositories.UserRepository.
func (m *mockUserRepo) GetUserById(issuer string) (entities.User, error) {
	return m.getUserByIdFunc(issuer)
}

// List implements repositories.UserRepository.
func (m *mockUserRepo) List() ([]entities.User, error) {
	panic("unimplemented")
}

func (m *mockUserRepo) Save(user entities.User) (entities.User, error) {
	return m.saveFunc(user)
}

func TestSaveUser(t *testing.T) {
	t.Run("save user success", func(t *testing.T) {
		repo := &mockUserRepo{
			saveFunc: func(user entities.User) (entities.User, error) {
				user = entities.User{
					Name:     "name",
					Email:    "email",
					Password: []byte("$2a$14$wZ2Ix8QEY5QaUvvLSZ5rGe7oOW/BW6Bd5iI2ovEVzNqPqcE9n3QmG"),
				}

				return user, nil
			},
		}

		service := usercase.NewUserUseCase(repo)

		_, err := service.Save(
			entities.User{
				Name:     "name",
				Email:    "email",
				Password: []byte("$2a$14$wZ2Ix8QEY5QaUvvLSZ5rGe7oOW/BW6Bd5iI2ovEVzNqPqcE9n3QmG"),
			},
		)
		assert.NoError(t, err)
	})

}
