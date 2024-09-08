package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"r_keeper/configs"
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/pkg/controllers"
	"r_keeper/server"
	"syscall"
)

func main() {
	if err := configs.ReadSettings(); err != nil {
		panic(err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	if err := db.ConnectToDB(); err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	mainServer := new(server.Server)
	go func() {
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.RunRouts()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	// Ожидание сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершения программ\n")
}
