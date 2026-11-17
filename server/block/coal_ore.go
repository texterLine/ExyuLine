// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// CoalOre is a common ore.
type CoalOre struct {
	solid
	bassDrum

	// Type is the type of coal ore.
	Type OreType
}

// BreakInfo ...
func (c CoalOre) BreakInfo() BreakInfo {
	return newBreakInfo(c.Type.Hardness(), pickaxeHarvestable, pickaxeEffective, oreDrops(item.Coal{}, c)).withXPDropRange(0, 2).withBlastResistance(3)
}

// SmeltInfo ...
func (CoalOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.Coal{}, 1), 0.1)
}

// EncodeItem ...
func (c CoalOre) EncodeItem() (name string, meta int16) {
	return "minecraft:" + c.Type.Prefix() + "coal_ore", 0
}

// EncodeBlock ...
func (c CoalOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + c.Type.Prefix() + "coal_ore", nil

}
