package main

import app "poker-game/internal"

func main() {
	newApp := app.GetApp()
	newApp.Run("0.0.0.0:8080")
}
