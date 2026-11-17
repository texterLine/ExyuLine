// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// SmoothBasalt is a decorative solid block obtained by smelting basalt.
type SmoothBasalt struct {
	solid
	bassDrum
}

// EncodeBlock ...
func (SmoothBasalt) EncodeBlock() (string, map[string]any) {
	return "minecraft:smooth_basalt", nil
}

// EncodeItem ...
func (SmoothBasalt) EncodeItem() (name string, meta int16) {
	return "minecraft:smooth_basalt", 0
}

// BreakInfo ...
func (s SmoothBasalt) BreakInfo() BreakInfo {
	return newBreakInfo(1.25, pickaxeHarvestable, pickaxeEffective, oneOf(s)).withBlastResistance(4.2)
}
