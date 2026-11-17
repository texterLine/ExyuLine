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

// RawIron is a raw metal block equivalent to nine raw iron.
type RawIron struct {
	solid
	bassDrum
}

// BreakInfo ...
func (r RawIron) BreakInfo() BreakInfo {
	return newBreakInfo(5, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, pickaxeEffective, oneOf(r)).withBlastResistance(6)
}

// EncodeItem ...
func (RawIron) EncodeItem() (name string, meta int16) {
	return "minecraft:raw_iron_block", 0
}

// EncodeBlock ...
func (RawIron) EncodeBlock() (string, map[string]any) {
	return "minecraft:raw_iron_block", nil
}
