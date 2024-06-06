package users

type RegisterRequestDTO struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	AuthProvider string `json:"auth_provider"`
}

type LoginRequestDTO struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	AuthProvider string `json:"auth_provider"`
}

type ResponseDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserRole string `json:"user_role"`
}

func (r *ResponseDTO) ToString() map[string]interface{} {
	return map[string]interface{}{
		"id":        r.ID,
		"email":     r.Email,
		"user_role": r.UserRole,
	}
}

func CreateResponseDTOFromModel(model Model) ResponseDTO {
	return ResponseDTO{
		ID:       model.ID.String(),
		Email:    model.Email,
		UserRole: model.UserRole,
	}
}
