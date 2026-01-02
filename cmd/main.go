package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"chat-app/internal/api/handler"
	"chat-app/internal/api/router"
	"chat-app/internal/api/ws"
	"chat-app/internal/repository"
	"chat-app/internal/services"
	"chat-app/internal/usecase"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	srv := buildServer()

	_ = ws.NewHub() // websocket

	wg.Add(2)
	go func() { // fix
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error server")
		}
	}()

	go func() { // fix
		defer wg.Done()
	}()

	<-ctx.Done()

	wg.Wait()
}

func buildServer() http.Server {
	msgRepo := repository.NewMsgRepo() // add postgresql
	roomRepo := repository.NewRoomRepo()
	userRepo := repository.NewUserRepo()

	msgSrv := services.NewMsgSvc(msgRepo)
	roomSrv := services.NewRoomSvc(roomRepo)
	userSrv := services.NewUserSvc(userRepo)

	us := usecase.NewUsecase(msgSrv, roomSrv, userSrv)

	handler := handler.NewHandler(us)

	router := router.NewRouter(handler)

	return http.Server{
		Addr:    "8080:", // add get from config
		Handler: router,
	}
}
