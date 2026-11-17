// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// NetherQuartzOre is ore found in the Nether.
type NetherQuartzOre struct {
	solid
	bassDrum
}

// BreakInfo ...
func (q NetherQuartzOre) BreakInfo() BreakInfo {
	return newBreakInfo(3, pickaxeHarvestable, pickaxeEffective, oreDrops(item.NetherQuartz{}, q)).withXPDropRange(0, 3)
}

// SmeltInfo ...
func (NetherQuartzOre) SmeltInfo() item.SmeltInfo {
	return newOreSmeltInfo(item.NewStack(item.NetherQuartz{}, 1), 0.2)
}

// EncodeItem ...
func (NetherQuartzOre) EncodeItem() (name string, meta int16) {
	return "minecraft:quartz_ore", 0
}

// EncodeBlock ...
func (NetherQuartzOre) EncodeBlock() (string, map[string]any) {
	return "minecraft:quartz_ore", nil
}
