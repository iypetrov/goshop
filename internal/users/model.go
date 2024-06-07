package users

import (
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"strings"
	"time"
)

type Model struct {
	ID           uuid.UUID
	Email        string
	Nickname     string
	Password     string
	AuthProvider string
	UserRole     string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

func CreateModelFromEntity(entity Entity) Model {
	return Model{
		ID:           entity.ID,
		Email:        entity.Email,
		Nickname:     entity.Nickname,
		Password:     entity.Password,
		AuthProvider: entity.AuthProvider,
		UserRole:     entity.UserRole,
		CreatedAt:    entity.CreatedAt,
		ModifiedAt:   entity.ModifiedAt,
	}
}

func CreateModelFromRegisterRequestDTO(requestDTO RegisterRequestDTO) Model {
	var model Model
	model.ID = uuid.New()
	model.Email = requestDTO.Email
	model.Nickname = requestDTO.Nickname
	model.Password = requestDTO.Password
	model.AuthProvider = ConvertToAuthProvider(requestDTO.AuthProvider)
	model.UserRole = CLIENT
	model.CreatedAt = time.Now()
	model.ModifiedAt = time.Now()
	return model
}

func CreateModelFromLoginRequestDTO(requestDTO LoginRequestDTO) Model {
	var model Model
	model.Email = requestDTO.Email
	model.Nickname = requestDTO.Nickname
	model.Password = requestDTO.Password
	model.AuthProvider = ConvertToAuthProvider(requestDTO.AuthProvider)
	return model
}

func (m *Model) Validate() error {
	var errors []string

	_, err := mail.ParseAddress(m.Email)
	if err != nil {
		errors = append(errors, "Email is not valid")
	}

	if m.AuthProvider == INVALID {
		errors = append(errors, "Auth provider is no valid")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}
