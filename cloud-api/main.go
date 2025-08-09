package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(dbFlushMiddleware())
	r.POST("/api/login", loginHandler)
	r.POST("/api/secret/register", registerHandler)
	r.GET("/api/users", usersHandler)
	r.GET("/api/ping", pingHandler)

	r.GET("/api/infrastructures", getAllInfrastructuresHandler)
	r.POST("/api/infrastructure", setInfrastructuresHandler)
	r.POST("/api/infrastructures", getInfrastructuresHandler)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
