package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserStorage_SaveUser(t *testing.T) {
	var userStorage UserStorage

	user1 := User{}
	user1.ID = 1
	user1.FirstName = "Francisco"
	user1.LastName = "Sanchez"
	user1.Age = 99

	userStorage.AddUser(user1)

	assert.Equal(t, 1, userStorage.Count())
}

func TestUserStorage_SearchUser(t *testing.T) {
	var userStorage UserStorage

	user1 := User{}
	user1.ID = 1
	user1.FirstName = "Francisco"
	user1.LastName = "Sanchez"
	user1.Age = 99

	userStorage.AddUser(user1)

	assert.Equal(t, user1, userStorage.FindUserByID(user1.ID))
}

func TestUserStorage_EmptySearchUser(t *testing.T) {
	var userStorage UserStorage

	assert.Equal(t, 0, userStorage.Count())
	assert.Equal(t, User{}, userStorage.FindUserByID(1))
}
