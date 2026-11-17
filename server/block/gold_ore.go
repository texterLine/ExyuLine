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

// GoldOre is a rare mineral block found underground.
type GoldOre struct {
	solid
	bassDrum

	// Type is the type of gold ore.
	Type OreType
}

// BreakInfo ...
func (g GoldOre) BreakInfo() BreakInfo {
	return newBreakInfo(g.Type.Hardness(), func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierIron.HarvestLevel
	}, pickaxeEffective, oreDrops(item.RawGold{}, g)).withBlastResistance(3)
}

// SmeltInfo ...
func (GoldOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.GoldIngot{}, 1), 1)
}

// EncodeItem ...
func (g GoldOre) EncodeItem() (name string, meta int16) {
	return "minecraft:" + g.Type.Prefix() + "gold_ore", 0
}

// EncodeBlock ...
func (g GoldOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + g.Type.Prefix() + "gold_ore", nil
}
