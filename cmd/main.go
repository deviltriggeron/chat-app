package main

import (
	"chat-app/internal/api/handler"
	"chat-app/internal/api/router"
	"chat-app/internal/api/ws"
	"chat-app/internal/config"
	"chat-app/internal/infrastructure"
	"chat-app/internal/repository"
	"chat-app/internal/services"
	"chat-app/internal/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
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
		fmt.Printf("Listen and serve: %s\n", "") // get port from config.env
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
	DBConfig := config.GetDBConfig()

	db := infrastructure.InitDB(DBConfig)

	msgRepo := repository.NewMsgRepo(db)
	roomRepo := repository.NewRoomRepo(db)
	userRepo := repository.NewUserRepo(db)

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
