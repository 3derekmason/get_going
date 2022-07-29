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
	router.POST("/entries", postEntries)
	router.Run("localhost:5000")
}

// getEntries responds with the list of all entries as JSON.
func getEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

// postEntries adds an entry from JSON received in the request body.
func postEntries(c *gin.Context) {
	var newEntry entry

	// Call BindJSON to bind the received JSON to
	// newEntry.
	if err := c.BindJSON(&newEntry); err != nil {
			return
	}

	// Add the new entry to the slice.
	events = append(events, newEntry)
	c.IndentedJSON(http.StatusCreated, newEntry)
}

// curl http://localhost:5000/entries \
// --include \
// --header "Content-Type: application/json" \
// --request "POST" \
// --data '{"id": "4","event": "Jog","distance": 2,"time": 15}'