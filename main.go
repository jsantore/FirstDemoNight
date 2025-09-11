package main

import (
	"fmt"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type firstGame struct {
	player       *ebiten.Image
	xloc         int
	yloc         int
	speed        int
	score        int
	allTreasures []treasure
}

type treasure struct {
	pict *ebiten.Image
	xLoc float64
	yLoc float64
}

func (f *firstGame) Update() error {
	f.xloc += f.speed
	if f.xloc > (1000-f.player.Bounds().Dx()) || f.xloc < 0 {
		f.speed = -f.speed
	}
	return nil
}

func (f *firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Mediumseagreen)
	drawOps := &ebiten.DrawImageOptions{}
	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(float64(f.xloc), float64(f.yloc))
	screen.DrawImage(f.player, drawOps)
	for _, treasure := range f.allTreasures {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(treasure.xLoc, treasure.yLoc)
		screen.DrawImage(treasure.pict, drawOps)
	}
}

func (f *firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerPict, _, err := ebitenutil.NewImageFromFile("ship.png")
	coinImage, _, err := ebitenutil.NewImageFromFile("coins.png")
	treasureSlice := make([]treasure, 0)
	for i := 0; i < 10; i++ {
		treasureSlice = append(treasureSlice, NewTreasure(950, 950, coinImage))
	}
	ourGame := firstGame{
		player:       playerPict,
		speed:        3,
		xloc:         400,
		yloc:         200,
		allTreasures: treasureSlice,
	} //we will use the zero value for now
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}

func NewTreasure(maxX, maxY int, image *ebiten.Image) treasure {
	return treasure{
		pict: image,
		xLoc: float64(rand.Intn(maxX)),
		yLoc: float64(rand.Intn(maxY)),
	}
}
