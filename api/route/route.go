package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, mux *http.ServeMux) {
	NewSignUpRouter(env, timeout, db, mux)
	NewLoginRouter(env, timeout, db, mux)

	NewProfileRouter(env, timeout, db, mux)

}
