package main

import (
	"fmt"
	"log"
	"os"
	"proj-erp-auth/controllers"
	"proj-erp-auth/database"
	"proj-erp-auth/middlewares"
	"proj-erp-auth/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	fmt.Println("User:", os.Getenv("DB_USER"))
	fmt.Println("Pass:", os.Getenv("DB_PASSWORD"))

	// Cria e configura o router corretamente
	router := gin.Default()

	// Conecta com o DB Postgres
	db, err := database.ConnectGORM()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	router.Use(middlewares.AuthMiddleware())

	router.GET("/user/:userId", controllers.GetUser(userRepository))

	router.Run(":8080")
}
