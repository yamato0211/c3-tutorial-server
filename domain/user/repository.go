package userrepository

import (
	"context"

	usermodel "github.com/yamato0211/c3-tutorial-server/domain/user/model"
)

type UserRepository interface {
	UpdateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error)
	GetUserByFirebaseID(ctx context.Context, firebaseID string) (*usermodel.User, error)
}
