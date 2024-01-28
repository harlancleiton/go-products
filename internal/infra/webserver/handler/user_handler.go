package handler

import (
	"encoding/json"
	"net/http"

	"github.com/harlancleiton/go-products/internal/dto"
	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/harlancleiton/go-products/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userDb database.UserInterface
}

func NewUserHandler(userDb database.UserInterface) *UserHandler {
	return &UserHandler{userDb: userDb}
}

// Create user godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body dto.CreateUserInput true "Create User"
// @Success 201 {object} dto.UserOutput "Created"
// @Failure 400 {object} Error "Bad Request"
// @Failure 500 {object} Error "Internal Server Error"
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.userDb.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	output := dto.NewUserOutput(user.ID.String(), user.Name, user.Email)
	json.NewEncoder(w).Encode(output)
}
