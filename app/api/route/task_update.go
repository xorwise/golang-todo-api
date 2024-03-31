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

func NewUpdateTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	mw := middleware.JWTMiddleware{
		Secret: env.AccessTokenSecret, Repository: repository.NewUserRepository(db),
	}
	tc := controller.UpdateTaskController{
		TaskUsecase: usecase.NewTaskUsecase(repository.NewTaskRepository(db), timeout),
		Env:         env,
	}

	mux.Handle("PUT /task/{id}", mw.LoginRequired(http.HandlerFunc(tc.Update)))
}
