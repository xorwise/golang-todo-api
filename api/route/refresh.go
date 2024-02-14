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

func NewRefreshRoute(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	rc := controller.RefreshController{
		RefreshUsecase: usecase.NewRefreshUsecase(repository.NewUserRepository(db), timeout),
		Env:            env,
	}

	mux.HandleFunc("POST /refresh", rc.Refresh)
}
