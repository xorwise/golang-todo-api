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

func NewCreateTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	mw := middleware.JWTMiddleware{Secret: env.AccessTokenSecret, Repository: repository.NewUserRepository(db)}
	tr := repository.NewTaskRepository(db)
	tc := controller.CreateTaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
		Env:         env,
	}

	mux.Handle("POST /task", mw.LoginRequired(http.HandlerFunc(tc.Create)))
}
