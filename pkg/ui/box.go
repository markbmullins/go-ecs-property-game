package ui

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

type Box struct {
	X, Y          float64
	Width, Height float64
	Padding       float64
	BgColor       color.Color
	TextColor     color.Color
	FontFace      font.Face
}

// Draw renders the box with a background color and centered multi-line text.
func (b *Box) Draw(screen *ebiten.Image, content string) {
	boxImage := ebiten.NewImage(int(b.Width), int(b.Height))
	boxImage.Fill(b.BgColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(boxImage, op)

	// Center multi-line text inside box.
	lines := strings.Split(content, "\n")
	startY := int(b.Y) + int(b.Padding)
	for i, line := range lines {
		x := int(b.X) + int((b.Width-float64(len(line)*8))/2) // Approximation for character width
		y := startY + i*16 // 16 is an approximation for line height
		ebitenutil.DebugPrintAt(screen, line, x, y)
	}
}

func (b *Box) IsClicked(screenX, screenY int) bool {
	return float64(screenX) >= b.X && float64(screenX) <= b.X+b.Width &&
		float64(screenY) >= b.Y && float64(screenY) <= b.Y+b.Height
}