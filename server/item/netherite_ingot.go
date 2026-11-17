// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/sandertv/gophertunnel/minecraft/text"

// NetheriteIngot is a rare mineral crafted with 4 pieces of netherite scrap and 4 gold ingots.
type NetheriteIngot struct{}

// EncodeItem ...
func (NetheriteIngot) EncodeItem() (name string, meta int16) {
	return "minecraft:netherite_ingot", 0
}

// TrimMaterial ...
func (NetheriteIngot) TrimMaterial() string {
	return "netherite"
}

// MaterialColour ...
func (NetheriteIngot) MaterialColour() string {
	return text.Netherite
}

// PayableForBeacon ...
func (NetheriteIngot) PayableForBeacon() bool {
	return true
}
