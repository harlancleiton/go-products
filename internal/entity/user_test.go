package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@mail.com", "123456")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@mail.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}
