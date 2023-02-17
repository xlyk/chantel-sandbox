package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	// start serving prometheus metrics
	ServeMetrics()

	// configure gin runtime mode
	ginMode := gin.ReleaseMode
	if os.Getenv("DEBUG") == "true" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// create gin router
	r := gin.Default()

	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// ping endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// date endpoint
	r.GET("/date", func(c *gin.Context) {
		DateRequestCounter.Inc()

		c.JSON(http.StatusOK, gin.H{
			"message": time.Now().Format(time.UnixDate),
		})
	})

	// start gin webserver
	if err := r.Run(); err != nil {
		panic(err)
	}
}
