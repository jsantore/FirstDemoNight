package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	_ "image/png"
)

type firstGame struct {
	player *ebiten.Image
	xloc   int
	yloc   int
	score  int
}

func (f firstGame) Update() error {
	return nil
}

func (f firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Mediumseagreen)
}

func (f firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerPict, _, err := ebitenutil.NewImageFromFile("player.png")
	ourGame := firstGame{
		player: playerPict,
		xloc:   400,
		yloc:   200,
	} //we will use the zero value for now
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
