// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/sandertv/gophertunnel/minecraft/text"

// ResinBrick is an item used to create resin bricks. It can also be used as a
// smithing ingredient, giving orange details to pieces of armour.
type ResinBrick struct{}

// EncodeItem ...
func (ResinBrick) EncodeItem() (name string, meta int16) {
	return "minecraft:resin_brick", 0
}

// TrimMaterial ...
func (ResinBrick) TrimMaterial() string {
	return "resin"
}

// MaterialColour ...
func (ResinBrick) MaterialColour() string {
	return text.Resin
}
