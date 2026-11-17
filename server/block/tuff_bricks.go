// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// TuffBricks are a decorational variant of Tuff that can be crafted or found naturally in Trial Chambers.
type TuffBricks struct {
	solid
	bassDrum

	// Chiseled specifies if the tuff bricks are chiseled.
	Chiseled bool
}

// BreakInfo ...
func (t TuffBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(t)).withBlastResistance(6)
}

// EncodeItem ...
func (t TuffBricks) EncodeItem() (name string, meta int16) {
	if t.Chiseled {
		return "minecraft:chiseled_tuff_bricks", 0
	}
	return "minecraft:tuff_bricks", 0
}

// EncodeBlock ...
func (t TuffBricks) EncodeBlock() (string, map[string]any) {
	if t.Chiseled {
		return "minecraft:chiseled_tuff_bricks", nil
	}
	return "minecraft:tuff_bricks", nil
}
