// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// EndStone is a block found in The End.
type EndStone struct {
	solid
	bassDrum
}

// BreakInfo ...
func (e EndStone) BreakInfo() BreakInfo {
	return newBreakInfo(3, pickaxeHarvestable, pickaxeEffective, oneOf(e)).withBlastResistance(9)
}

// EncodeItem ...
func (EndStone) EncodeItem() (name string, meta int16) {
	return "minecraft:end_stone", 0
}

// EncodeBlock ...
func (EndStone) EncodeBlock() (string, map[string]any) {
	return "minecraft:end_stone", nil
}
