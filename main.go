package main

import (
	"sportify_services/config"
	"sportify_services/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectionDatabase()
	config.Migration()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Ganti dengan origin frontend kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Metode yang diizinkan
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Header yang diizinkan
	}))

	v1 := r.Group("/v1")

	routers.Index(v1)

	r.Run()
}