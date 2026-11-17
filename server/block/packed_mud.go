// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// PackedMud is a block crafted from mud and wheat. It is used to create mud bricks.
type PackedMud struct {
	solid
}

// BreakInfo ...
func (p PackedMud) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, nothingEffective, oneOf(p)).withBlastResistance(3)
}

// EncodeItem ...
func (PackedMud) EncodeItem() (name string, meta int16) {
	return "minecraft:packed_mud", 0
}

// EncodeBlock ...
func (PackedMud) EncodeBlock() (string, map[string]any) {
	return "minecraft:packed_mud", nil
}
