package database

import (
	"testing"

	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser(
		"John Doe",
		"johndoe@mail.com",
		"123456",
	)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.Find(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.True(t, userFound.ComparePassword("123456"))
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser(
		"John Doe",
		"johndoe@mail.com",
		"123456",
	)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}
