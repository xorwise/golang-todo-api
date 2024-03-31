package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type TaskWebsocketController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *TaskWebsocketController) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := tc.Env.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		return
	}

	tc.Env.Mu.Lock()
	tc.Env.ClientChannels[userID] = make(chan string)
	tc.Env.Mu.Unlock()

	defer func() {
		tc.Env.Mu.Lock()
		close(tc.Env.ClientChannels[userID])
		delete(tc.Env.ClientChannels, userID)
		tc.Env.Mu.Unlock()
	}()

	go func() {

		for msg := range tc.Env.ClientChannels[userID] {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
		}
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			return
		}
	}
}
