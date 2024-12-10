package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к базе данных
	database.ConnectDatabase()

	// Инициализация сервиса
	bookService := &services.BookService{}

	// Создаем экземпляр маршрутизатора
	r := gin.Default()

	// Настройка маршрутов
	routes.SetupBookRoutes(r, bookService)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
