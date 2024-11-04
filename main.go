// main.go
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/markbmullins/city-developer/pkg/constants"
	"github.com/markbmullins/city-developer/pkg/game"
)

func main() {
	game := game.NewGame()

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Property Management Prototype")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
