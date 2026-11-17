// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// TropicalFish is a food item that cannot be cooked.
type TropicalFish struct {
	defaultFood
}

// Consume ...
func (TropicalFish) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(1, 0.2)
	return Stack{}
}

// EncodeItem ...
func (TropicalFish) EncodeItem() (name string, meta int16) {
	return "minecraft:tropical_fish", 0
}
