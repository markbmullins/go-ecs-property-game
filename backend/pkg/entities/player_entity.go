package entities

import (
	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

/** Creates a player entity in the game.
 * A player entity has the following components:
 * Nameable: The name of the player.
 * Funds: The current funds available to the player.
 */
func CreatePlayer(
	id int,
	name string,
	initialFunds float64,
) *ecs.Entity {
	player := ecs.NewEntity("Player", id)

	ecs.AddComponent(player, &components.Nameable{Name: name})
	ecs.AddComponent(player, &components.Funds{Amount: initialFunds})

	return player
}
