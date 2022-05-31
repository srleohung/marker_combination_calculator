package main

import (
	"marker_combination_calculator_backend/task"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var t *task.Task = task.New()

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "200 OK",
	})
}

// Implement an API to calculate the total number of possible combination.
func NewTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": t.Start(),
	})
}

// Implement an API to show the list of unique marker configuration.
func GetResult(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": t.GetResult(),
	})
}

// Implement an API so that task progress can be read.
func GetProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": t.GetProgress(),
	})
}

// Implement an API to cancel the running task.
func CancelTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": t.Cancel(),
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", ping)
	r.GET("/new_task", NewTask)
	r.GET("/get_result", GetResult)
	r.GET("/get_progress", GetProgress)
	r.GET("/cancel_task", CancelTask)
	r.Run()
}
