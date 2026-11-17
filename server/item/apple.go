// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world"
)

// Apple is a food item that can be eaten by the player.
type Apple struct {
	defaultFood
}

// Consume ...
func (a Apple) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(4, 2.4)
	return Stack{}
}

// CompostChance ...
func (Apple) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (a Apple) EncodeItem() (name string, meta int16) {
	return "minecraft:apple", 0
}
