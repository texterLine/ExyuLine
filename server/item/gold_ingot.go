// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/sandertv/gophertunnel/minecraft/text"

// GoldIngot is a metal ingot melted from raw gold or obtained from loot chests.
type GoldIngot struct{}

// EncodeItem ...
func (GoldIngot) EncodeItem() (name string, meta int16) {
	return "minecraft:gold_ingot", 0
}

// TrimMaterial ...
func (GoldIngot) TrimMaterial() string {
	return "gold"
}

// MaterialColour ...
func (GoldIngot) MaterialColour() string {
	return text.Gold
}

// PayableForBeacon ...
func (GoldIngot) PayableForBeacon() bool {
	return true
}
