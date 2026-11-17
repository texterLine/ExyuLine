// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// DeepslateTiles are a tiled variant of deepslate and can spawn in ancient cities.
type DeepslateTiles struct {
	solid
	bassDrum

	// Cracked specifies if the deepslate tiles is its cracked variant.
	Cracked bool
}

// BreakInfo ...
func (d DeepslateTiles) BreakInfo() BreakInfo {
	return newBreakInfo(3.5, pickaxeHarvestable, pickaxeEffective, oneOf(d)).withBlastResistance(6)
}

// SmeltInfo ...
func (d DeepslateTiles) SmeltInfo() item.SmeltInfo {
	if d.Cracked {
		return item.SmeltInfo{}
	}
	return newSmeltInfo(item.NewStack(DeepslateTiles{Cracked: true}, 1), 0.1)
}

// EncodeItem ...
func (d DeepslateTiles) EncodeItem() (name string, meta int16) {
	if d.Cracked {
		return "minecraft:cracked_deepslate_tiles", 0
	}
	return "minecraft:deepslate_tiles", 0
}

// EncodeBlock ...
func (d DeepslateTiles) EncodeBlock() (string, map[string]any) {
	if d.Cracked {
		return "minecraft:cracked_deepslate_tiles", nil
	}
	return "minecraft:deepslate_tiles", nil
}
