package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserOutput(id, name, email string) *UserOutput {
	return &UserOutput{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
