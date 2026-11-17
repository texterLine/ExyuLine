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

// AncientDebris is a rare ore found within The Nether.
type AncientDebris struct {
	solid
}

// BreakInfo ...
func (a AncientDebris) BreakInfo() BreakInfo {
	return newBreakInfo(30, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierDiamond.HarvestLevel
	}, pickaxeEffective, oneOf(a)).withBlastResistance(1200)
}

// SmeltInfo ...
func (AncientDebris) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.NetheriteScrap{}, 1), 2)
}

// EncodeItem ...
func (AncientDebris) EncodeItem() (name string, meta int16) {
	return "minecraft:ancient_debris", 0
}

// EncodeBlock ...
func (AncientDebris) EncodeBlock() (string, map[string]any) {
	return "minecraft:ancient_debris", nil
}
