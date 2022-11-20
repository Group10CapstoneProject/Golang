package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

// UpdateSessionID implements AuthRepository
func (r *authRepositoryImpl) UpdateSessionID(userId uint, sessionId uuid.UUID, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", userId).Update("session_id", sessionId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrUserNotFound
	}
	return nil
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		db: database,
	}
}
