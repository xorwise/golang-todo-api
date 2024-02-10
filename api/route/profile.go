package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/api/controller"
	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/repository"
	"github.com/xorwise/golang-todo-api/usecase"
	"gorm.io/gorm"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {

	mw := middleware.JWTMiddleware{
		Secret: env.AccessTokenSecret,
	}
	ur := repository.NewUserRepository(db)
	sc := controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
		Env:            env,
	}

	mux.Handle("/profile", mw.LoginRequired(http.HandlerFunc(sc.GetProfile)))
}
