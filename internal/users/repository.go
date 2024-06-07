package users

import (
	"context"
	"fmt"
	"github.com/iypetrov/goshop/internal/common"
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

// UpdateEntity Only can update nickname and modified_at
func (r *Repository) UpdateEntity(entity Entity) (Entity, error) {
	tx, errTx, cl := common.CreateTx(r.ctx, r.conn)
	if errTx != nil {
		slog.Error(errTx.Error())
		return Entity{}, errTx
	}
	defer cl()

	var exists bool
	err := tx.QueryRow(
		r.ctx,
		`SELECT EXISTS(SELECT 1 FROM "user" WHERE id=$1)`,
		entity.ID,
	).Scan(&exists)
	if err != nil {
		slog.Error(err.Error())
		return Entity{}, err
	}
	if !exists {
		slog.Error(fmt.Sprintf("user with id %s does not exist", entity.ID))
		return Entity{}, err
	}

	var updatedEntity Entity
	err = tx.QueryRow(
		r.ctx,
		`UPDATE "user" 
		SET nickname=$1, modified_at=$2 
		WHERE id=$3 
		RETURNING id, email, password, auth_provider, user_role, created_at, modified_at`,
		entity.Nickname,
		entity.ModifiedAt,
		entity.ID,
	).Scan(
		&updatedEntity.ID,
		&updatedEntity.Email,
		&updatedEntity.Nickname,
		&updatedEntity.Password,
		&updatedEntity.AuthProvider,
		&updatedEntity.UserRole,
		&updatedEntity.CreatedAt,
		&updatedEntity.ModifiedAt,
	)
	if err != nil {
		slog.Error(err.Error())
		return Entity{}, err
	}

	return updatedEntity, nil
}

func (r *Repository) DeleteEntity(id uuid.UUID) error {
	_, err := r.conn.Exec(
		r.ctx,
		`DELETE FROM "user" WHERE id=$1`,
		id,
	)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
