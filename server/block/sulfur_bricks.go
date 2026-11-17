// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// SulfurBricks is a decorative variant of Sulfur.
type SulfurBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (s SulfurBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(s)).withBlastResistance(6)
}

// EncodeItem ...
func (SulfurBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:sulfur_bricks", 0
}

// EncodeBlock ...
func (SulfurBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:sulfur_bricks", nil
}
