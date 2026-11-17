// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/item/potion"

// Arrow is used as ammunition for bows, crossbows, and dispensers. Arrows can be modified to
// imbue status effects on players and mobs.
type Arrow struct {
	// Tip is the potion effect that is tipped on the arrow.
	Tip potion.Potion
}

// EncodeItem ...
func (a Arrow) EncodeItem() (name string, meta int16) {
	if tip := a.Tip.Uint8(); tip > 4 {
		return "minecraft:arrow", int16(tip + 1)
	}
	return "minecraft:arrow", 0
}

// OffHand ...
func (Arrow) OffHand() bool {
	return true
}
