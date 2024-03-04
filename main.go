package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/routes"
	"github.com/gin-contrib/cors"
)

//NOTE 1 framework for rest api: go get -u github.com/gin-gonic/gin
//NOTE 2 use the MongoDB driver: go get go.mongodb.org/mongo-driver/mongo
//NOTE 4 run "go get github.com/golang-jwt/jwt/v5" to get the jwt package
//NOTE 5 run "go get golang.org/x/crypto/bcrypt" to get the bcrypt package
//NOTE 6 run "go get github.com/gin-contrib/cors" to get the cors package

func main() {
	db.InitDB()
	
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	routes.RegisterRoutes(server)

	server.Run(":8080")
}