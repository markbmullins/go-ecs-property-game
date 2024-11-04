// ui/ui.go
package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/markbmullins/city-developer/pkg/constants"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/utils"
	"golang.org/x/image/font"
)

// DrawUI renders the user interface elements.
func DrawUI(screen *ebiten.Image, playerFunds float64, gameDate string, isPaused bool, fontFace font.Face) {
    fundsText := fmt.Sprintf("Funds: $%.2f", playerFunds)
    dateText := fmt.Sprintf("Date: %s", gameDate)

	// Define reusable Box for text
	fundsBox := Box{
		X:         10,
		Y:         20,
		Width:     100,
		Height:    30,
		Padding:   5,
		BgColor:   color.White,
		TextColor: color.Black,
		FontFace:  fontFace,
	}
	fundsBox.Draw(screen, fundsText)

	dateBox := Box{
		X:         float64(constants.ScreenWidth - 150),
		Y:         20,
		Width:     140,
		Height:    30,
		Padding:   5,
		BgColor:   color.White,
		TextColor: color.Black,
		FontFace:  fontFace,
	}
	dateBox.Draw(screen, dateText)

	// Define Pause/Resume Button using Button struct
	pauseText := map[bool]string{true: "Resume", false: "Pause"}[isPaused]
	pauseButton := Button{
		Box: Box{
			X:         constants.ScreenWidth - 90,
			Y:         constants.ScreenHeight - 40,
			Width:     constants.ButtonWidth,
			Height:    constants.ButtonHeight,
			Padding:   constants.Padding,
			BgColor:   color.RGBA{200, 200, 200, 255},
			TextColor: color.Black,
			FontFace:  fontFace,
		},
		Content: pauseText,
		OnClick: func() { isPaused = !isPaused },
	}
	pauseButton.Draw(screen)
}
// DrawTileInfo renders information about a specific tile.
func DrawTileInfo(screen *ebiten.Image, tile *models.Tile, fontFace font.Face) {
    // Position for the popup
    x := tile.Position.X * constants.TileSize
    y := tile.Position.Y * constants.TileSize
    popupX := x
    popupY := y - 150
    if popupY < 0 {
        popupY = y + constants.TileSize
    }
    if popupX+140 > constants.ScreenWidth {
        popupX = constants.ScreenWidth - 140
    }

    // Restore full popup details
    var infoText string
    if tile.Property != nil {
        ownershipStatus := "Available"
        if tile.Property.Owned {
            ownershipStatus = "Owned"
        }
        cost := tile.LandValue * 100
        infoText = fmt.Sprintf(
            "Name: %s\nAddress: %s\nType: %s\nSubtype: %s\nStatus: %s\nLevel: %d\nIncome: $%.2f\nPurchase Cost: $%.2f\nUpgrade Cost: $%.2f",
            tile.Property.Name, tile.Property.Address, tile.Property.Type, tile.Property.Subtype,
            ownershipStatus, tile.Property.Level, tile.Property.BaseIncome, cost, tile.Property.UpgradeCost)
    } else {
        cost := tile.LandValue * 100
        infoText = fmt.Sprintf("Land Value: $%.2f\nPurchase Cost: $%.2f", tile.LandValue*100, cost)
    }

    // Split the text into lines for individual rendering
    lines := utils.SplitLines(infoText)
    infoHeight := len(lines)*15 + 50

    // Draw popup background
    infoRect := ebiten.NewImage(140, infoHeight)
    infoRect.Fill(color.RGBA{255, 255, 255, 220})
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(popupX), float64(popupY))
    screen.DrawImage(infoRect, op)

    // Draw each line of text with calculated vertical spacing
    for i, line := range lines {
        xPos := popupX + 5
        yPos := popupY + 15 + i*15
        text.Draw(screen, line, fontFace, xPos, yPos, color.Black)
    }

    // Draw action button (Buy or Upgrade)
    buttonWidth := constants.ButtonWidth
    buttonHeight := constants.ButtonHeight
    buttonY := float64(popupY) + float64(infoHeight) - float64(buttonHeight) - 5

    if tile.Property != nil && !tile.Property.Owned {
        // Draw "Buy" button
        purchaseButtonRect := ebiten.NewImage(buttonWidth, buttonHeight)
        purchaseButtonRect.Fill(color.RGBA{0, 200, 0, 255})
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(popupX+5), buttonY)
        screen.DrawImage(purchaseButtonRect, op)

        // Draw "Buy" text centered
        text.Draw(screen, "Buy", fontFace, int(popupX+20), int(buttonY+15), color.Black)
    } else if tile.Property != nil && tile.Property.Owned {
        // Draw "Upgrade" button
        upgradeButtonRect := ebiten.NewImage(buttonWidth, buttonHeight)
        upgradeButtonRect.Fill(color.RGBA{0, 0, 200, 255})
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(popupX+5), buttonY)
        screen.DrawImage(upgradeButtonRect, op)

        // Draw "Upgrade" text centered
        text.Draw(screen, "Upgrade", fontFace, int(popupX+8), int(buttonY+15), color.White)
    }
}
