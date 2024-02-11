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

func NewUpdateUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {

	ur := repository.NewUserRepository(db)
	mw := middleware.JWTMiddleware{
		Secret:     env.AccessTokenSecret,
		Repository: repository.NewUserRepository(db),
	}
	sc := controller.UpdateUserController{
		UpdateUserUsecase: usecase.NewUpdateUserUsecase(ur, timeout),
		Env:               env,
	}

	mux.Handle("PUT /profile", mw.LoginRequired(http.HandlerFunc(sc.Update)))

}
