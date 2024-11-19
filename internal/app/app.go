package app

import (
	"L0/config"
	"L0/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	port string
}

func NewServer() *Server {
	return &Server{
		echo: echo.New(),
		port: ":8080",
	}
}

func (s *Server) ListenAndServe(orderHand echo.HandlerFunc, allOrdersHand echo.HandlerFunc) error {
	s.echo.GET("/orders/:order", orderHand)
	s.echo.GET("/orders/get", allOrdersHand)

	return s.echo.Start(s.port)
}

func Run(cfg *config.Config) {

	conn, err := database.Connect(&cfg.DB)
	if err != nil {
		log.Fatalf("error connecting to PostgreSQL: %v", err)
	}
	defer conn.Close()
	db := database.NewDB(conn)

	err = db.CreateTable()
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}


	httpServer := NewServer()
	apiController := controller.NewOrderController(cache)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	go func() {
		err := httpServer.ListenAndServe(apiController.GetOrder, apiController.GetAllOrder)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server: %v", err)
		}
	}()

}
