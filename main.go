package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)
	router.GET("/entries/:id", getEntryByID)
	router.POST("/entries", postEntries)
	router.Run("localhost:5000")
}

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

// getEntryByID locates the entry whose ID value matches the id
// parameter sent by the client, then returns that entry as a response.
func getEntryByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of entries, looking for
	// an entry whose ID value matches the parameter.
	for _, a := range events {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "entry not found"})
}