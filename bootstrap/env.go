package bootstrap

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type Env struct {
	DBUser             string
	DBPass             string
	DBHost             string
	DBPort             string
	DBName             string
	AccessTokenSecret  string
	AccessTokenExpiry  int
	RefreshTokenSecret string
	RefreshTokenExpiry int
	ContextTimeout     time.Duration
	Upgrader           websocket.Upgrader
	ClientChannels     map[uint]chan string
	Mu                 sync.Mutex
}

func NewEnv() *Env {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := &Env{
		DBUser:             os.Getenv("DB_USER"),
		DBPass:             os.Getenv("DB_PASSWORD"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBName:             os.Getenv("DB_NAME"),
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenExpiry:  168,
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenExpiry: 672,
		ContextTimeout:     10 * time.Second,
		ClientChannels:     make(map[uint]chan string),
		Mu:                 sync.Mutex{},
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	return env

}
