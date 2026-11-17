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

// IronOre is a mineral block found underground.
type IronOre struct {
	solid
	bassDrum

	// Type is the type of iron ore.
	Type OreType
}

// BreakInfo ...
func (i IronOre) BreakInfo() BreakInfo {
	return newBreakInfo(i.Type.Hardness(), func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, pickaxeEffective, oreDrops(item.RawIron{}, i)).withBlastResistance(3)
}

// SmeltInfo ...
func (IronOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.IronIngot{}, 1), 0.7)
}

// EncodeItem ...
func (i IronOre) EncodeItem() (name string, meta int16) {
	return "minecraft:" + i.Type.Prefix() + "iron_ore", 0
}

// EncodeBlock ...
func (i IronOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + i.Type.Prefix() + "iron_ore", nil
}
