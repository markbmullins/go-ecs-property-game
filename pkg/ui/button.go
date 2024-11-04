// button.go
package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	Box
	Content string
	OnClick func()
}

// Draw renders the button with text centered
func (b *Button) Draw(screen *ebiten.Image) {
	b.Box.Draw(screen, b.Content)
}

// HandleClick executes OnClick if button is clicked
func (b *Button) HandleClick(screenX, screenY int) {
	if b.IsClicked(screenX, screenY) && b.OnClick != nil {
		b.OnClick()
	}
}
