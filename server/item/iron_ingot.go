// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/sandertv/gophertunnel/minecraft/text"

// IronIngot is a metal ingot melted from raw iron or obtained from loot chests.
type IronIngot struct{}

// EncodeItem ...
func (IronIngot) EncodeItem() (name string, meta int16) {
	return "minecraft:iron_ingot", 0
}

// TrimMaterial ...
func (IronIngot) TrimMaterial() string {
	return "iron"
}

// MaterialColour ...
func (IronIngot) MaterialColour() string {
	return text.Iron
}

// PayableForBeacon ...
func (IronIngot) PayableForBeacon() bool {
	return true
}
