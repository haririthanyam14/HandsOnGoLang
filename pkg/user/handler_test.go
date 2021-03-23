package user_test

import (
	"HandsOnGoLang/pkg/user"
	"HandsOnGoLang/pkg/user/mocks"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestAddUserHandlerSuccess(t *testing.T) {
	u := user.UserRequest{
		FirstName: "bruce",
		LastName:  "wayne",
	}

	us := new(mocks.Service)
	newUser, _ := user.NewUser(u)
	us.On("CreateUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*user.User")).Return(newUser, nil)
	expectedBody := `{"UserID":0,"Message":"Successfully inserted"}`

	testAddUserHandler(t, expectedBody, us, &u)
}

func testAddUserHandler(t *testing.T, expectedBody string, service user.Service, userDetails *user.UserRequest) {
	userReqJSON, err := json.Marshal(userDetails)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userReqJSON))
	require.NoError(t, err)

	res := httptest.NewRecorder()

	uh := user.NewUserHandler(service)

	(uh.CreateUser)(res, req)
	userResp, err := json.Marshal(res)
	log.Printf("user respone %s", userResp)
	assert.Equal(t, expectedBody, res.Body.String())
}
