package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/harlancleiton/go-products/internal/dto"
	"github.com/harlancleiton/go-products/internal/infra/database"
)

type AuthHandler struct {
	jwt       *jwtauth.JWTAuth
	expiresIn int
	userDb    *database.User
}

func NewAuthHandler(jwt *jwtauth.JWTAuth, expiresIn int, userDb *database.User) *AuthHandler {
	return &AuthHandler{jwt: jwt, expiresIn: expiresIn, userDb: userDb}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.CredentialsInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.userDb.FindByEmail(input.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ComparePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println(u)

	_, token, err := h.jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Duration(h.expiresIn) * time.Second).Unix(),
	})

	fmt.Println(token)
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}
