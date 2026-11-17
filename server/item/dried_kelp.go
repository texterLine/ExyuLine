// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// DriedKelp is a food item that can be quickly eaten by the player.
type DriedKelp struct{}

// AlwaysConsumable ...
func (DriedKelp) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (DriedKelp) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration / 2
}

// Consume ...
func (DriedKelp) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(1, 0.2)
	return Stack{}
}

// CompostChance ...
func (DriedKelp) CompostChance() float64 {
	return 0.3
}

// EncodeItem ...
func (DriedKelp) EncodeItem() (name string, meta int16) {
	return "minecraft:dried_kelp", 0
}
