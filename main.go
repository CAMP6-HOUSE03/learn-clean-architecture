package main

import (
	"clean-arc/handler"
	"clean-arc/model"
	"clean-arc/repository"
	"clean-arc/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// clean arc -> independent
// contract interface -> error juga independent
// tidak package serving server

// fungsi middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("username")
		password := c.GetHeader("password")

		if username != "afis" || password != "1234" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrResponse("unauthorized"))
			return
		}

		c.Next()
	}
}

func main() {
	// wiring process

	// setup repository
	repo := repository.NewRepository()

	// setup service
	svc := service.NewService(repo)

	//setup handler
	handler := handler.NewHandler(svc)

	// repo -> service -> handler
	//

	// bikin engine / gin rpute

	r := gin.Default()

	// http.HandleFunc("/calculate", handler.HandleCalculator)

	r.Use(Auth())

	r.POST("/calculate", handler.HandleCalculator)
	r.GET("/history", handler.GetHistoryCalculator)

	r.Run() // http.ListenAndServe(":8080", nil)

	// engine, gin router
}

// aku mau bikin 1 handler dengan logic sebagai berikut
// calculator, method POST
// ngirim data berupa json
// 3 => a, b, operator
// nyimpen history operasi
// a, b, operasi, hasilnya, tanggal
// hasilnya akan dikembalikan ke response
