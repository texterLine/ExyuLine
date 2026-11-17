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

// Diamond is a block which can only be gained by crafting it.
type Diamond struct {
	solid
}

// BreakInfo ...
func (d Diamond) BreakInfo() BreakInfo {
	return newBreakInfo(5, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierIron.HarvestLevel
	}, pickaxeEffective, oneOf(d)).withBlastResistance(6)
}

// PowersBeacon ...
func (Diamond) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (Diamond) EncodeItem() (name string, meta int16) {
	return "minecraft:diamond_block", 0
}

// EncodeBlock ...
func (Diamond) EncodeBlock() (string, map[string]any) {
	return "minecraft:diamond_block", nil
}
