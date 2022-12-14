package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (m *AuthRepositoryMock) UpdateSessionID(userId *int, sessionId *uuid.UUID, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
