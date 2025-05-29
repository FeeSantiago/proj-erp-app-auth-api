package main

import (
	"log"
	"proj-erp-auth/controllers"
	"proj-erp-auth/database"
	"proj-erp-auth/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	//Config logger
	router := gin.New()
	router.Use(gin.Logger())

	//Connecta com o DB Postgres
	db, err := database.ConnectGORM()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	userRepo := repository.NewUserRepository(db)

	r := gin.Default()

	// Passa o repo para o handler
	r.GET("/user/:userId", controllers.GetUser(userRepo))

	// Outros endpoints...

	r.Run(":8080")
}
