package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/api/controller"
	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/repository"
	"github.com/xorwise/golang-todo-api/usecase"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)

func NewTaskWebsocketRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	mw := middleware.JWTMiddleware{
		Secret: env.AccessTokenSecret, Repository: repository.NewUserRepository(db),
	}
	tc := controller.TaskWebsocketController{
		TaskUsecase: usecase.NewTaskUsecase(repository.NewTaskRepository(db), timeout),
		Env:         env,
	}

	mux.Handle("/ws", mw.LoginRequired(websocket.Handler(tc.HandleConnection)))
}
