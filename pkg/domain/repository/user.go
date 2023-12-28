package repository

import (
	"context"

	"github.com/yamato0211/c3-tutorial-server/pkg/domain/entity"
)

type UserRepository interface {
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByFirebaseID(ctx context.Context, firebaseID string) (*entity.User, error)
}
