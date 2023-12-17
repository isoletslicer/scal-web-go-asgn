package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(ctx *gin.Context) {
		var data Json

		randSource := rand.NewSource(time.Now().UnixNano())
		random := rand.New(randSource)

		data.Status.Water, data.Status.Wind = random.Intn(100), random.Intn(100)

		var waterStatus string
		if data.Status.Water > 8 {
			waterStatus = "bahaya"
		} else if data.Status.Water >= 6 {
			waterStatus = "siaga"
		} else {
			waterStatus = "aman"
		}

		var windStatus string
		if data.Status.Wind > 15 {
			windStatus = "bahaya"
		} else if data.Status.Wind >= 7 {
			windStatus = "siaga"
		} else {
			windStatus = "aman"
		}

		fmt.Println("{")
		fmt.Printf("\"water\":%d,\n", data.Status.Water)
		fmt.Printf("\"wind\":%d\n", data.Status.Wind)
		fmt.Println("}")
		fmt.Println("status water:", waterStatus)
		fmt.Println("status wind:", windStatus)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Water":       data.Status.Water,
			"WaterStatus": waterStatus,
			"Wind":        data.Status.Wind,
			"WindStatus":  windStatus,
		})
	})

	router.Run()
}
