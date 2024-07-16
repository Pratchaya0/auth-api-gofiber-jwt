package validators

import (
	"github.com/stretchr/testify/mock"
)

type MockUserHandler struct {
	mock.Mock
}

func (m *MockUserHandler) Login() error {
	args := m.Called()
	return args.Error(0)
}

// func TestLoginHandler(t *testing.T) {
// 	mockService := new(MockUserHandler)
// 	handler := handlers.NewAuthHandler(mockService)
// }
