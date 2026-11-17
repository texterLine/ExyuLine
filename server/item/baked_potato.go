// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// BakedPotato is a food item that can be eaten by the player.
type BakedPotato struct {
	defaultFood
}

// Consume ...
func (BakedPotato) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(5, 6)
	return Stack{}
}

// CompostChance ...
func (BakedPotato) CompostChance() float64 {
	return 0.85
}

// EncodeItem ...
func (BakedPotato) EncodeItem() (name string, meta int16) {
	return "minecraft:baked_potato", 0
}
