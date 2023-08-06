package baseServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func health(hr healthReady) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ok := hr.Health(ctx)
		if !ok {
			c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"status": "DOWN"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"status": "UP"})
	}
}

func ready(hr healthReady) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		statuses := hr.Ready(ctx)
		httpStatus := http.StatusOK
		for _, err := range statuses {
			if err != nil {
				httpStatus = http.StatusBadGateway
				break
			}
		}
		c.IndentedJSON(httpStatus, pingsToJSON(statuses))
		return
	}
}

func pingsToJSON(statuses map[string]error) gin.H {

	response := gin.H{}

	for service, err := range statuses {
		response[service] = err == nil
	}

	return response
}
