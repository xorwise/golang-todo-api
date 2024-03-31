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

func NewFetchTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	mw := middleware.JWTMiddleware{
		Secret: env.AccessTokenSecret, Repository: repository.NewUserRepository(db),
	}
	repo := repository.NewTaskRepository(db)
	tc := controller.FetchTaskController{
		TaskUsecase: usecase.NewTaskUsecase(repo, timeout),
		Env:         env,
	}

	mux.Handle("GET /tasks/{user_id}", mw.LoginRequired(http.HandlerFunc(tc.Fetch)))
}
