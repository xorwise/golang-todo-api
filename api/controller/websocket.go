package controller

import (
	"fmt"
	"net/http"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
	"golang.org/x/net/websocket"
)

type TaskWebsocketController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *TaskWebsocketController) HandleConnection(ws *websocket.Conn) {
	fmt.Println("handle connection")
	userID, ok := ws.Request().Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		ws.WriteClose(http.StatusUnauthorized)
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
			if err := websocket.Message.Send(ws, msg); err != nil {
				return
			}
		}
	}()

	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			return
		}
	}
}
