// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// Cobblestone is a common block, obtained from mining stone.
type Cobblestone struct {
	solid
	bassDrum

	// Mossy specifies if the cobblestone is mossy. This variant of cobblestone is typically found in
	// dungeons or in small clusters in the giant tree taiga biome.
	Mossy bool
}

// BreakInfo ...
func (c Cobblestone) BreakInfo() BreakInfo {
	return newBreakInfo(2, pickaxeHarvestable, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// SmeltInfo ...
func (Cobblestone) SmeltInfo() item.SmeltInfo {
	return newSmeltInfo(item.NewStack(Stone{}, 1), 0.1)
}

// RepairsStoneTools ...
func (c Cobblestone) RepairsStoneTools() bool {
	return !c.Mossy
}

// EncodeItem ...
func (c Cobblestone) EncodeItem() (name string, meta int16) {
	if c.Mossy {
		return "minecraft:mossy_cobblestone", 0
	}
	return "minecraft:cobblestone", 0
}

// EncodeBlock ...
func (c Cobblestone) EncodeBlock() (string, map[string]any) {
	if c.Mossy {
		return "minecraft:mossy_cobblestone", nil
	}
	return "minecraft:cobblestone", nil
}
