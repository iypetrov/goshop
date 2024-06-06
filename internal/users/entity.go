package users

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	AuthProvider string    `db:"auth_provider"`
	UserRole     string    `db:"user_role"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
}

func CreateEntityFromModel(model Model) Entity {
	return Entity{
		ID:           model.ID,
		Email:        model.Email,
		Password:     model.Password,
		AuthProvider: model.AuthProvider,
		UserRole:     model.UserRole,
		CreatedAt:    model.CreatedAt,
		ModifiedAt:   model.ModifiedAt,
	}
}
