package login

type LoginDto struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
