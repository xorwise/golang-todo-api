package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/api/controller"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/repository"
	"github.com/xorwise/golang-todo-api/usecase"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	ur := repository.NewUserRepository(db)
	sc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	mux.HandleFunc("POST /login", sc.Login)

}
