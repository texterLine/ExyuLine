// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// MagmaCream is an item used in brewing to create potions of Fire Resistance, and to build magma blocks.
type MagmaCream struct{}

// EncodeItem ...
func (m MagmaCream) EncodeItem() (name string, meta int16) {
	return "minecraft:magma_cream", 0
}
