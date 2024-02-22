package main

import (
	"github.com/forwork0529/for_open_graph_protokol/internal/api_rest"
	"github.com/forwork0529/for_open_graph_protokol/pkg/config"
	"github.com/forwork0529/for_open_graph_protokol/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// определяем переменные окружения
	err := config.LoadFromFile(".env")

	// инициализируем логер
	logger.New(config.Values.LoggerLevel)

	// проверка ошибки подгрузки env var
	if err != nil {
		logger.Fatal(err)
	}

	// server
	server := api_rest.NewServer()

	// Run http server
	api_rest.RunGin(server, config.Values.APIEndpoint)

	// ожидание сигнала прерывания и запуск завершения внутренних процессов
	awaitQuitSignal()

	// функции непосредственно завершающие открытые соединения:

	// ожидание завершения процессов остановки
	waitingAndReturn()
}

func awaitQuitSignal() {
	logger.Infof("Microservice started. Working until a quit signal is received...")
	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Infof("awaitQuitSignal(): Stopping server...")
}

func waitingAndReturn() {
	time.Sleep(time.Second * 2)
	logger.Info("waitingAndReturn(): The stop timer is completed")
}
