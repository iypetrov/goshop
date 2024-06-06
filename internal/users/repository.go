package users

import (
	"context"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

var (
	repository *Repository
	oncerep    sync.Once
)

func NewRepository(ctx context.Context, conn *pgxpool.Pool) {
	oncerep.Do(func() {
		repository = &Repository{
			ctx:  ctx,
			conn: conn,
		}
	})
}

func GetRepository() *Repository {
	return repository
}

func (r *Repository) GetEntityByID(id uuid.UUID) (Entity, error) {
	var entity Entity
	err := r.conn.QueryRow(
		r.ctx,
		`SELECT id, email, password, auth_provider, user_role,  created_at FROM "user" WHERE id = $1`,
		id,
	).Scan(
		&entity.ID,
		&entity.Email,
		&entity.Password,
		&entity.AuthProvider,
		&entity.UserRole,
		&entity.CreatedAt,
	)
	if err != nil {
		slog.Error(err.Error())
		return Entity{}, err
	}

	return entity, nil
}

func (r *Repository) GetEntityByEmail(email string, authProvider string) (Entity, error) {
	var entity Entity
	err := r.conn.QueryRow(
		r.ctx,
		`SELECT id, email, password, auth_provider, user_role, created_at, modified_at FROM "user" WHERE email = $1 AND auth_provider = $2`,
		email,
		authProvider,
	).Scan(
		&entity.ID,
		&entity.Email,
		&entity.Password,
		&entity.AuthProvider,
		&entity.UserRole,
		&entity.CreatedAt,
		&entity.ModifiedAt,
	)
	if err != nil {
		slog.Error(err.Error())
		return Entity{}, err
	}

	return entity, nil
}

func (r *Repository) CreateEntity(entity Entity) (Entity, error) {
	_, err := r.conn.Exec(
		r.ctx,
		`INSERT INTO "user" (id, email, password, auth_provider, user_role, created_at, modified_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.AuthProvider,
		entity.UserRole,
		entity.CreatedAt,
		entity.ModifiedAt,
	)
	if err != nil {
		slog.Error(err.Error())
		return Entity{}, err
	}

	return entity, nil
}
