package users

type RegisterRequestDTO struct {
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	Password     string `json:"password"`
	AuthProvider string `json:"auth_provider"`
}

type LoginRequestDTO struct {
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	Password     string `json:"password"`
	AuthProvider string `json:"auth_provider"`
}

type ResponseDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	UserRole string `json:"user_role"`
}

func (r *ResponseDTO) ToString() map[string]interface{} {
	return map[string]interface{}{
		"id":        r.ID,
		"email":     r.Email,
		"nickname":  r.Nickname,
		"user_role": r.UserRole,
	}
}

func CreateResponseDTOFromModel(model Model) ResponseDTO {
	return ResponseDTO{
		ID:       model.ID.String(),
		Email:    model.Email,
		Nickname: model.Nickname,
		UserRole: model.UserRole,
	}
}
