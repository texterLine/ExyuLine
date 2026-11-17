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

// Beetroot is a food and dye ingredient.
type Beetroot struct {
	defaultFood
}

// Consume ...
func (b Beetroot) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(1, 1.2)
	return Stack{}
}

// CompostChance ...
func (Beetroot) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (b Beetroot) EncodeItem() (name string, meta int16) {
	return "minecraft:beetroot", 0
}
