package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/*
func getBallotByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range ballots {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
 }
*/

func setCurrentBallot(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("type of id is: ", reflect.TypeOf(id))

	if err := dbSetCurrrentBallotByID(db, id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, "")
}

func getCurrentBallot(c *gin.Context) {
	fmt.Println("Dancers - getting ballot")
	bal, err := dbGetCurrentBallot(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, bal)
	fmt.Println("Dancers - Returned ballot: ", bal)
}

func getAllBallots(c *gin.Context) {
	allBallots, err := dbGetAllBallots(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, allBallots)
}

func getClearVotes(c *gin.Context) {
	dbClearVotes(db)
	c.IndentedJSON(http.StatusOK, "")
}

func getPopulateDB(c *gin.Context) {
	populateDB(db)
	c.IndentedJSON(http.StatusOK, "")
}

func getFadeState(c *gin.Context) {
	fad, err := dbGetFadeState(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, fad)
	fmt.Println("Current fade state: ", fad)
}

func toggleFadeState(c *gin.Context) {
	dbToggleState(db)
	c.IndentedJSON(http.StatusOK, "Fade state toggled")
}

func fadeTrue(c *gin.Context) {
	dbFadeTrue(db)
	c.IndentedJSON(http.StatusOK, "Fade state set to true")
}

func fadeFalse(c *gin.Context) {
	dbFadeFalse(db)
	c.IndentedJSON(http.StatusOK, "Fade state set to false")
}

func getExportDB(c *gin.Context) {
	dbExportDB(db)
	c.IndentedJSON(http.StatusOK, "")
}

// TODO: are we using this? Delete if no
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Resolve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TODO: are we using this? Delete if no
func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

// TODO: are we using this? Delete if no
func voteSocket(c *gin.Context) {
	wshandler(c.Writer, c.Request)
	return
	// c.String(http.StatusOK, "Hi, WebSocket!")
	// return
}

func addVote(c *gin.Context) {
	var newVote voteStruct
	// fmt.Println(c)

	// Call BindJSON to bind the received JSON to voteStruct
	if err := c.BindJSON(&newVote); err != nil {
		return
	}

	bal, err := dbVote(db, newVote)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not vote")
		return
	}
	c.IndentedJSON(http.StatusOK, bal)
	// c.IndentedJSON(http.StatusOK, bal)
}
