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

func NewSignUpRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignUpController{
		SignUpUsecase: usecase.NewSignUpUsecase(ur, timeout),
		Env:           env,
	}
	mux.HandleFunc("/signup", sc.SignUp)
}
