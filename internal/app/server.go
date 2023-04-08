package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/internal/handlers"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(port string) error {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/rest/substr/find", handlers.HandleFindSubStr).Methods("POST")
	router.HandleFunc("/rest/email/check", handlers.HandleFindEmail).Methods("POST")
	router.HandleFunc("/rest/iin/check", handlers.HandleFindIIN).Methods("POST")
	router.HandleFunc("/rest/self/find/{substr}", handlers.HandleFindSelf).Methods("GET")

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           "localhost:" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	// канал quit, на который регистрируется обработчик для сигнала os.Interrupt (обычно это сигнал Ctrl-C в консоли).
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	// Главная горутина блокируется на чтении из канала quit, ожидая получения сигнала os.Interrupt.
	<-quit

	// Если получен сигнал os.Interrupt, создается новый контекст ctx с помощью функции context.WithTimeout(),
	// который дает 5 секунд на завершение работы сервера.
	// Затем метод Shutdown() вызывается с этим контекстом, чтобы закрыть HTTP-сервер.
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
