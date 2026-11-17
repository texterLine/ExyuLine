// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
)

// Melon is a fruit block that grows from melon stems.
type Melon struct {
	solid
}

// BreakInfo ...
func (m Melon) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, axeEffective, discreteDrops(item.MelonSlice{}, m, 3, 7, 9))
}

// CompostChance ...
func (Melon) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (Melon) EncodeItem() (name string, meta int16) {
	return "minecraft:melon_block", 0
}

// EncodeBlock ...
func (Melon) EncodeBlock() (string, map[string]any) {
	return "minecraft:melon_block", nil
}
