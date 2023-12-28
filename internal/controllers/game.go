package controllers

import (
	"github.com/gin-gonic/gin"
)

func JoinAGame(c *gin.Context) {
	// Look for existing games if there is seat
	// in a existing game join the user to that game

	// If not create a new game and the user to it as waiting for
	// other players

	// Once joined the table if a hand is running the player wont be allowed
	// to play that hand and will start at the begining of the other
	// hand its frontends responsibilty to start a hand

}
