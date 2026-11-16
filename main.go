package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkServiceHealth(service string) {
	resp, err := http.Get(service + "/health")
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("Service `%v` is down. \n", service)
	}
}

func roundRobin(services []string, counter *int) string {
	service := services[*counter]
	*counter++

	if *counter >= len(services) {
		*counter = 0
	}

	return service
}

func callTask(c *gin.Context, service string) {
	resp, err := http.Post(fmt.Sprintf("%s/task", service),
		"application/json", nil,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "service unavailable",
		})
		return
	}
	c.Data(resp.StatusCode, "application/json", body)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	DeployedServices := []string{"http://localhost:8000", "http://localhost:8001", "http://localhost:8002", "http://localhost:8003"}

	counter := 0
	patterns := make(map[string]bool)
	patterns["roundrobin"] = true

	fmt.Println("----------------------------")
	fmt.Println("Deployed Services")
	for idx, service := range DeployedServices {
		fmt.Printf("%v, %v \n", idx, service)
	}
	fmt.Println("----------------------------")

	for _, service := range DeployedServices {
		checkServiceHealth(service)
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Load balancer root route",
		})
	})

	router.GET("/message", func(ctx *gin.Context) {
		if patterns["roundrobin"] {
			service := roundRobin(DeployedServices, &counter)
			fmt.Println("********************************************************************")
			fmt.Printf("Forwarding task to %v \n", service)
			fmt.Printf("Current counter %v \n", counter)
			fmt.Println("********************************************************************")
			callTask(ctx, service)
		}
	})

	router.Run(":8080")

}
