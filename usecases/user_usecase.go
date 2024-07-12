package usecases

import (
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/repositories"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (useCase *UserUseCase) GetUserById(id string) (entities.User, error) {
	return useCase.userRepo.GetUserById(id)
}

func (useCase *UserUseCase) GetUserByEmail(email string) (entities.User, error) {
	return useCase.userRepo.GetUserByEmail(email)
}

func (useCase *UserUseCase) Save(user entities.User) (entities.User, error) {
	return useCase.userRepo.Save(user)
}

func (useCase *UserUseCase) List() ([]entities.User, error) {
	return useCase.userRepo.List()
}
