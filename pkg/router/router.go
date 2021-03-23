package router

import (
	hello "HandsOnGoLang/pkg/hello"
	"HandsOnGoLang/pkg/user"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Routers(lgr *zap.Logger, us user.Service) http.Handler {

	r := mux.NewRouter()
	uh := user.NewUserHandler(us)

	r.HandleFunc("/hello", hello.Hello)
	r.HandleFunc("/headers", hello.Headers)
	r.HandleFunc("/headers/error", hello.HelloError)

	r.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

	return r

}
