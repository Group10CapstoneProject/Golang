package repository

import (
	"context"

	"github.com/google/uuid"
)

type AuthRepository interface {
	UpdateSessionID(userId uint, sessionId uuid.UUID, ctx context.Context) error
}
