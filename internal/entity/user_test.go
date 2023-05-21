package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "password")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "password")

	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("passwd"))
	assert.NotEqual(t, "password", user.Password)
}
