package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	user, err := NewUser("Erick Farias", "erick.farias@gmail.com", "123456")

	assert.Nil(t, err)
	assert.Equal(t, user.Username, "Erick Farias")
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, user.Password, "1234567")
	assert.True(t, user.ValidatePassword("123456"))

}

func TestUser_ValidatePassword(t *testing.T) {

	user, err := NewUser("Erick Farias", "erick.farias@gmail.com", "123456")

	validPassword := user.ValidatePassword("123456")
	invalidPassword := user.ValidatePassword("1234567")

	assert.Nil(t, err)
	assert.True(t, validPassword)
	assert.False(t, invalidPassword)
	assert.NotEqual(t, user.Password, "1234567")

}
