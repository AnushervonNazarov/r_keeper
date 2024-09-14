package main

import (
	"context"
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
	"time"

	"github.com/joho/godotenv"
)

// @title R_KEEPER API
// @version 1.0
// @description R-Keeper: Restaurant management software.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
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

	// Закрытие соединения с БД, если необходимо
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение с БД успешно закрыто")

	// Используем контекст с тайм-аутом для завершения работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")
}
