// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// CinnabarBricks is a decorative variant of Cinnabar.
type CinnabarBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (c CinnabarBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// EncodeItem ...
func (CinnabarBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:cinnabar_bricks", 0
}

// EncodeBlock ...
func (CinnabarBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:cinnabar_bricks", nil
}
