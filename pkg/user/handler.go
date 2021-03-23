package user

import (
	"HandsOnGoLang/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	us Service
}

func NewUserHandler(userService Service) *UserHandler {
	return &UserHandler{
		us: userService,
	}
}

func (uh *UserHandler) CreateUser(resp http.ResponseWriter, req *http.Request) {
	var ur UserRequest
	err := utils.ParseRequest(req, &ur)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("parser error"))
		return
	}
	user, _ := NewUser(ur)
	log.Printf("ur.FullName  %s", user.FullName)
	ctx := req.Context()
	res, _ := uh.us.CreateUser(ctx, user)
	//res, _ := NewUserService(req.Context()).CreateUser(&user)
	userResp := UserResp{
		UserID:  res.ID,
		Message: "Successfully inserted",
	}
	result, _ := json.Marshal(&userResp)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(result)
}
