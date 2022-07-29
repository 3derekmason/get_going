package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// entry represents data about an event.
type entry struct {
	ID     string  `json:"id"`
	Event  string  `json:"event"`
	Distance float64  `json:"distance"`
	Time  float64 `json:"time"`
}



// events slice to seed record event data.
var events = []entry{
	{ID: "1", Event: "Jog", Distance: 1, Time: 8},
	{ID: "2", Event: "Hike", Distance: 3.5, Time: 70},
	{ID: "3", Event: "Walk", Distance: 2.25, Time: 30},
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)

	router.Run("localhost:5000")
}

// getEntries responds with the list of all entries as JSON.
func getEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}