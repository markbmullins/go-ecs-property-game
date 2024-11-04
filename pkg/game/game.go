// game/game.go
package game

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/markbmullins/city-developer/pkg/constants"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/ui"
	"github.com/markbmullins/city-developer/pkg/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/*.ttf
var fontFiles embed.FS

// Game struct represents the state of the game.
type Game struct {
    Grid              [][]*models.Tile
    PlayerFunds       float64
    IncomeTimer       float64
    EthicalChoiceMade bool
    Message           string
    HoveredTile       *models.Tile

    GameDate  time.Time
    GameSpeed float64
    IsPaused  bool

    FontFace font.Face
}

// loadFont initializes and returns a font.Face.
func loadFont() font.Face {
    // Read the font data from the embedded files
    fontData, err := fontFiles.ReadFile("fonts/Roboto-Regular.ttf")
    if err != nil {
        log.Fatalf("Failed to read font data: %v", err)
    }

    // Parse the font
    f, err := opentype.Parse(fontData)
    if err != nil {
        log.Fatalf("Failed to parse font: %v", err)
    }

    // Create a font face
    face, err := opentype.NewFace(f, &opentype.FaceOptions{
        Size:    14,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        log.Fatalf("Failed to create font face: %v", err)
    }

    return face
}

// NewGame initializes and returns a new Game instance.
func NewGame() *Game {
    fontFace := loadFont()

    g := &Game{
        PlayerFunds:       constants.InitialPlayerFunds,
        IncomeTimer:       constants.IncomeInterval,
        EthicalChoiceMade: false,
        Message:           "",
        GameDate:          time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
        GameSpeed:         1.0,
        IsPaused:          false,
        FontFace:          fontFace,
    }
    g.InitializeGrid()
    return g
}

// InitializeGrid sets up the game grid with tiles and properties.
func (g *Game) InitializeGrid() {
    cols := constants.ScreenWidth / constants.TileSize
    rows := constants.ScreenHeight / constants.TileSize
    g.Grid = make([][]*models.Tile, rows)
    for y := 0; y < rows; y++ {
        g.Grid[y] = make([]*models.Tile, cols)
        for x := 0; x < cols; x++ {
            tile := &models.Tile{
                Position:  models.Position{X: x, Y: y},
                LandValue: 1.0 + rand.Float64()*9.0,
            }
            propertyType, subtypes := randomPropertyType()
            subtype := subtypes[rand.Intn(len(subtypes))]
            name := fmt.Sprintf("%s %s", propertyType, subtype)
            address := fmt.Sprintf("%d Main St", y*cols+x+1)
            tile.Property = &models.Property{
                Position:    tile.Position,
                Level:       0,
                BaseIncome:  10.0,
                UpgradeCost: 200.0,
                Owned:       false,
                Name:        name,
                Address:     address,
                Type:        propertyType,
                Subtype:     subtype,
            }
            g.Grid[y][x] = tile
        }
    }
}
// DrawGrid draws the game grid and its tiles.
func (g *Game) DrawGrid(screen *ebiten.Image) {
    for y, row := range g.Grid {
        for x, tile := range row {
            // Determine the color of the tile
            color := tileColor(tile)
            if tile.Property != nil && tile.Property.Owned {
                color = propertyColor(tile.Property)
            }

            // Create an image for the tile
            tileImage := ebiten.NewImage(constants.TileSize, constants.TileSize)
            tileImage.Fill(color)

            // Set the position for drawing
            op := &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(x*constants.TileSize), float64(y*constants.TileSize))

            // Draw the tile onto the screen
            screen.DrawImage(tileImage, op)
        }
    }
}


// randomPropertyType randomly selects a property type and its subtypes.
func randomPropertyType() (string, []string) {
    types := []struct {
        Type     string
        Subtypes []string
    }{
        {"Residential", []string{"Single Family", "Multifamily", "Townhome"}},
        {"Commercial", []string{"Office", "Retail", "Industrial"}},
    }
    selectedType := types[rand.Intn(len(types))]
    return selectedType.Type, selectedType.Subtypes
}

// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    g.HandleInput()
    g.UpdateHoveredTile()

    if !g.IsPaused {
        elapsed := 1.0 / 60.0 * g.GameSpeed
        g.GameDate = g.GameDate.Add(time.Duration(elapsed * float64(time.Second)))
        g.IncomeTimer -= elapsed
        if g.IncomeTimer <= 0 {
            g.GenerateIncome()
            g.IncomeTimer = constants.IncomeInterval
        }
    }

    return nil
}

// Draw is called every frame (typically 60 times per second).
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.White)
    g.DrawGrid(screen)
    ui.DrawUI(
        screen,
        g.PlayerFunds,
        g.GameDate.Format("Jan 2, 2006"),
        g.IsPaused,
        g.FontFace,
    )
    if g.HoveredTile != nil {
        ui.DrawTileInfo(screen, g.HoveredTile, g.FontFace)
    }
    
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return constants.ScreenWidth, constants.ScreenHeight
}

// HandleInput processes user inputs.
func (g *Game) HandleInput() {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()

        buttonWidth := constants.ButtonWidth
        buttonHeight := constants.ButtonHeight
        buttonX := constants.ScreenWidth - buttonWidth - 10
        buttonY := constants.ScreenHeight - buttonHeight - 10
        if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
            g.IsPaused = !g.IsPaused
            g.Message = "Game " + map[bool]string{true: "paused", false: "resumed"}[g.IsPaused]
            return
        }

        for i, s := range []struct {
            Label string
            Speed float64
        }{
            {"Slow", 0.5},
            {"Normal", 1.0},
            {"Fast", 2.0},
        } {
            btnX := 10 + i*(constants.ButtonWidth+10)
            btnY := constants.ScreenHeight - constants.ButtonHeight - 10
            if x >= btnX && x <= btnX+constants.ButtonWidth && y >= btnY && y <= btnY+constants.ButtonHeight {
                g.GameSpeed = s.Speed
                g.Message = fmt.Sprintf("Game speed set to %s.", s.Label)
                return
            }
        }

        g.HandleClick(x, y)
    }

    if !g.EthicalChoiceMade {
        if ebiten.IsKeyPressed(ebiten.KeyB) {
            g.BribeOfficial()
            g.EthicalChoiceMade = true
            g.Message = "You chose to bribe the official."
        } else if ebiten.IsKeyPressed(ebiten.KeyI) {
            g.InvestInCommunity()
            g.EthicalChoiceMade = true
            g.Message = "You invested in the community."
        }
    }
}

// HandleClick processes clicks on the game grid.
func (g *Game) HandleClick(screenX, screenY int) {
    if g.HoveredTile != nil {
        x := g.HoveredTile.Position.X * constants.TileSize
        y := g.HoveredTile.Position.Y * constants.TileSize
        popupX := x
        popupY := y - 150
        if popupY < 0 {
            popupY = y + constants.TileSize
        }
        if popupX+140 > constants.ScreenWidth {
            popupX = constants.ScreenWidth - 140
        }

        var infoText string
        if g.HoveredTile.Property != nil {
            ownershipStatus := "Available"
            if g.HoveredTile.Property.Owned {
                ownershipStatus = "Owned"
            }
            cost := g.HoveredTile.LandValue * 100
            infoText = fmt.Sprintf(
                "Name: %s\nAddress: %s\nType: %s\nSubtype: %s\nStatus: %s\nLevel: %d\nIncome: $%.2f\nPurchase Cost: $%.2f\nUpgrade Cost: $%.2f",
                g.HoveredTile.Property.Name, g.HoveredTile.Property.Address, g.HoveredTile.Property.Type, g.HoveredTile.Property.Subtype,
                ownershipStatus, g.HoveredTile.Property.Level, g.HoveredTile.Property.BaseIncome, cost, g.HoveredTile.Property.UpgradeCost)
        } else {
            cost := g.HoveredTile.LandValue * 100
            infoText = fmt.Sprintf("Land Value: $%.2f\nPurchase Cost: $%.2f", g.HoveredTile.LandValue*100, cost)
        }
        lines := utils.SplitLines(infoText)
        infoHeight := len(lines)*15 + 50

        buttonWidth := constants.ButtonWidth
        buttonHeight := constants.ButtonHeight
        buttonX := float64(popupX + 5)
        buttonY := float64(popupY) + float64(infoHeight) - float64(buttonHeight) - 5

        if float64(screenX) >= buttonX && float64(screenX) <= buttonX+float64(buttonWidth) &&
            float64(screenY) >= buttonY && float64(screenY) <= buttonY+float64(buttonHeight) {
            if g.HoveredTile.Property != nil && !g.HoveredTile.Property.Owned {
                g.BuyProperty(g.HoveredTile)
            } else if g.HoveredTile.Property != nil && g.HoveredTile.Property.Owned {
                g.UpgradeProperty(g.HoveredTile.Property)
            }
            return
        }
    }
}

// BuyProperty handles purchasing a property.
func (g *Game) BuyProperty(tile *models.Tile) {
    cost := tile.Property.BaseIncome * 10 // Adjusted: Assuming Purchase Cost is BaseIncome * 10
    if g.PlayerFunds >= cost {
        g.PlayerFunds -= cost
        tile.Property.Owned = true
        tile.Property.Level = 1
        g.Message = "Property purchased."
    } else {
        g.Message = "Not enough funds to buy property."
    }
}

// UpgradeProperty handles upgrading a property.
func (g *Game) UpgradeProperty(p *models.Property) {
    if g.PlayerFunds >= p.UpgradeCost {
        g.PlayerFunds -= p.UpgradeCost
        p.Level++
        p.BaseIncome *= 1.5
        p.UpgradeCost *= 1.5
        g.IncreaseAdjacentLandValues(p.Position)
        g.Message = "Property upgraded."
    } else {
        g.Message = "Not enough funds to upgrade."
    }
}

// GenerateIncome calculates and adds income from owned properties.
func (g *Game) GenerateIncome() {
    totalIncome := 0.0
    for _, row := range g.Grid {
        for _, tile := range row {
            if tile.Property != nil && tile.Property.Owned {
                income := tile.Property.BaseIncome * float64(tile.Property.Level)
                income -= income * constants.TaxRate
                totalIncome += income
            }
        }
    }
    g.PlayerFunds += totalIncome
    g.Message = fmt.Sprintf("Income generated: $%.2f", totalIncome)
}

// IncreaseAdjacentLandValues increases land values adjacent to a given position.
func (g *Game) IncreaseAdjacentLandValues(pos models.Position) {
    directions := []models.Position{
        {X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0},
    }
    for _, dir := range directions {
        newX := pos.X + dir.X
        newY := pos.Y + dir.Y
        if newY >= 0 && newY < len(g.Grid) && newX >= 0 && newX < len(g.Grid[0]) {
            adjTile := g.Grid[newY][newX]
            adjTile.LandValue += 0.1
        }
    }
}

// BribeOfficial handles the bribe action.
func (g *Game) BribeOfficial() {
    bribeCost := 500.0
    if g.PlayerFunds >= bribeCost {
        g.PlayerFunds -= bribeCost
        for _, row := range g.Grid {
            for _, tile := range row {
                if tile.Property != nil && tile.Property.Owned {
                    tile.Property.UpgradeCost *= 0.8
                }
            }
        }
        g.Message = "Bribe successful. Upgrade costs reduced."
    } else {
        g.Message = "Not enough funds to bribe."
    }
}

// InvestInCommunity handles the invest action.
func (g *Game) InvestInCommunity() {
    investmentCost := 500.0
    if g.PlayerFunds >= investmentCost {
        g.PlayerFunds -= investmentCost
        for _, row := range g.Grid {
            for _, tile := range row {
                if tile.Property != nil && tile.Property.Owned {
                    tile.Property.BaseIncome *= 1.1
                }
            }
        }
        g.Message = "Investment successful. Income increased."
    } else {
        g.Message = "Not enough funds to invest."
    }
}

// UpdateHoveredTile updates the currently hovered tile based on cursor position.
func (g *Game) UpdateHoveredTile() {
    x, y := ebiten.CursorPosition()
    gridX := x / constants.TileSize
    gridY := y / constants.TileSize

    // Update HoveredTile based on cursor position within grid bounds
    if gridY >= 0 && gridY < len(g.Grid) && gridX >= 0 && gridX < len(g.Grid[0]) {
        g.HoveredTile = g.Grid[gridY][gridX]
    } else {
        g.HoveredTile = nil
        return
    }

    // Popup position and text setup for hovered tile
    if g.HoveredTile != nil {
        tileX := g.HoveredTile.Position.X * constants.TileSize
        tileY := g.HoveredTile.Position.Y * constants.TileSize
        popupX := tileX
        popupY := tileY - 150
        if popupY < 0 {
            popupY = tileY + constants.TileSize
        }
        if popupX+140 > constants.ScreenWidth {
            popupX = constants.ScreenWidth - 140
        }

        // Generate detailed info text
        var infoText string
        if g.HoveredTile.Property != nil {
            ownershipStatus := "Available"
            if g.HoveredTile.Property.Owned {
                ownershipStatus = "Owned"
            }
            cost := g.HoveredTile.LandValue * 100
            infoText = fmt.Sprintf(
                "Name: %s\nAddress: %s\nType: %s\nSubtype: %s\nStatus: %s\nLevel: %d\nIncome: $%.2f\nPurchase Cost: $%.2f\nUpgrade Cost: $%.2f",
                g.HoveredTile.Property.Name, g.HoveredTile.Property.Address, g.HoveredTile.Property.Type, g.HoveredTile.Property.Subtype,
                ownershipStatus, g.HoveredTile.Property.Level, g.HoveredTile.Property.BaseIncome, cost, g.HoveredTile.Property.UpgradeCost)
        } else {
            cost := g.HoveredTile.LandValue * 100
            infoText = fmt.Sprintf("Land Value: $%.2f\nPurchase Cost: $%.2f", g.HoveredTile.LandValue*100, cost)
        }

        // Determine popup dimensions
        lines := utils.SplitLines(infoText)
        infoHeight := len(lines)*15 + 50
        popupWidth := 140
        popupHeight := infoHeight

        // Prevent redundant updates if cursor is within popup bounds
        if float64(x) >= float64(popupX) && float64(x) <= float64(popupX)+float64(popupWidth) &&
            float64(y) >= float64(popupY) && float64(y) <= float64(popupY)+float64(popupHeight) {
            return
        }
    }
}

// tileColor determines the color of a tile based on its land value.
func tileColor(tile *models.Tile) color.Color {
    shade := uint8(200 - tile.LandValue*10)
    if shade < 100 {
        shade = 100
    }
    return color.RGBA{shade, shade, shade, 255}
}

// propertyColor determines the color of a property based on its level.
func propertyColor(p *models.Property) color.Color {
    shade := uint8(100 + p.Level*30)
    if shade > 255 {
        shade = 255
    }
    return color.RGBA{shade, 200, shade, 255}
}
