package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	// Public routes
	NewSignUpRouter(env, timeout, db, mux)
	NewLoginRouter(env, timeout, db, mux)

	// Protected routes
	NewProfileRouter(env, timeout, db, mux)
	NewUpdateUserRouter(env, timeout, db, mux)
	NewCreateTaskRouter(env, timeout, db, mux)
	NewFetchTaskRouter(env, timeout, db, mux)
	NewUpdateTaskRouter(env, timeout, db, mux)
	NewDeleteTaskRouter(env, timeout, db, mux)

	// Websocket routes
	NewTaskWebsocketRouter(env, timeout, db, mux)
}
