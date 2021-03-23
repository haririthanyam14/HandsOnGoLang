package user_test

import (
	"HandsOnGoLang/pkg/user"
	"HandsOnGoLang/pkg/user/mocks"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserSuccess(t *testing.T) {
	u := &user.User{
		FirstName: "bruce",
		LastName:  "wayne",
		FullName:  "bruce wayne",
	}

	userRepo := new(mocks.UserRepository)
	userRepo.On("AddUser", context.Background(), u).Return(nil)

	testAddUser(t, nil, u, userRepo)
}

func testAddUser(t *testing.T, expectedError error, userDetails *user.User, userRepo user.UserRepository) {
	us := user.NewUserService(userRepo)
	res, _ := us.CreateUser(context.Background(), userDetails)
	//_, err:= us.CreateUser(context.Background(), userDetails)
	log.Printf("inside service test")
	assert.Equal(t, "bruce", res.FirstName)
	//assert.Equal(t, expectedError, err)
}
