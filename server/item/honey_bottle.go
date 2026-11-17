// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"time"

	"github.com/texterLine/exyuline/server/entity/effect"
	"github.com/texterLine/exyuline/server/world"
)

// HoneyBottle is a food item obtained from full beehives or beenests using a glass bottle. Consuming it
// restores hunger, removes any active poison and returns an empty glass bottle.
type HoneyBottle struct{}

// MaxCount ...
func (HoneyBottle) MaxCount() int {
	return 16
}

// AlwaysConsumable ...
func (HoneyBottle) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (HoneyBottle) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration * 5 / 4
}

// Consume ...
func (HoneyBottle) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(6, 1.2)
	c.RemoveEffect(effect.Poison)
	return NewStack(GlassBottle{}, 1)
}

// EncodeItem ...
func (HoneyBottle) EncodeItem() (name string, meta int16) {
	return "minecraft:honey_bottle", 0
}
