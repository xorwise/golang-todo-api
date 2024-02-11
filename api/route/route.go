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

}
