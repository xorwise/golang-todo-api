package main

import (
	"log"
	"net/http"

	"github.com/xorwise/golang-todo-api/api/route"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/internal/worker"
)

func main() {
	env := bootstrap.NewEnv()

	db := bootstrap.NewDatabaseConnection(env)
	err := bootstrap.MigrateDatabase(db)
	if err != nil {
		log.Fatal(err)
	}
	defer bootstrap.CloseDatabaseConnection(db)

	go worker.CheckDeadlines(env, db)
	timeout := env.ContextTimeout

	mux := http.NewServeMux()

	route.Setup(env, timeout, db, mux)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
