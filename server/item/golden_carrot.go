// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// GoldenCarrot is a valuable food item and brewing ingredient. It provides the second most saturation in the game,
// behind Suspicious Stew crafted with either a Dandelion or Blue Orchid.
type GoldenCarrot struct {
	defaultFood
}

// Consume ...
func (GoldenCarrot) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(6, 14.4)
	return Stack{}
}

// EncodeItem ...
func (GoldenCarrot) EncodeItem() (name string, meta int16) {
	return "minecraft:golden_carrot", 0
}
