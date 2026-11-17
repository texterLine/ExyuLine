// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// GlisteringMelonSlice is an inedible item used for brewing potions of healing. It is also one of the many potion
// ingredients that can be used to make mundane potions.
type GlisteringMelonSlice struct{}

// EncodeItem ...
func (GlisteringMelonSlice) EncodeItem() (name string, meta int16) {
	return "minecraft:glistering_melon_slice", 0
}
