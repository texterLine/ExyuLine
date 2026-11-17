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

// Netherite is a precious mineral block made from 9 netherite ingots.
type Netherite struct {
	solid
	bassDrum
}

// BreakInfo ...
func (n Netherite) BreakInfo() BreakInfo {
	return newBreakInfo(50, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierDiamond.HarvestLevel
	}, pickaxeEffective, oneOf(n)).withBlastResistance(1200)
}

// PowersBeacon ...
func (Netherite) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (Netherite) EncodeItem() (name string, meta int16) {
	return "minecraft:netherite_block", 0
}

// EncodeBlock ...
func (Netherite) EncodeBlock() (string, map[string]any) {
	return "minecraft:netherite_block", nil
}
