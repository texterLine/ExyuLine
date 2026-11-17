// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// ResinBricks is a block crafted from resin brick.
type ResinBricks struct {
	solid
	bassDrum

	// Chiseled specifies if the resin bricks is its chiseled variant.
	Chiseled bool
}

// BreakInfo ...
func (r ResinBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(r)).withBlastResistance(6)
}

// EncodeItem ...
func (r ResinBricks) EncodeItem() (name string, meta int16) {
	if r.Chiseled {
		return "minecraft:chiseled_resin_bricks", 0
	}
	return "minecraft:resin_bricks", 0
}

// EncodeBlock ...
func (r ResinBricks) EncodeBlock() (string, map[string]any) {
	if r.Chiseled {
		return "minecraft:chiseled_resin_bricks", nil
	}
	return "minecraft:resin_bricks", nil
}
