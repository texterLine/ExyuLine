// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Dripstone is a rock block that allows pointed dripstone to grow beneath it.
type Dripstone struct {
	solid
	bassDrum
}

// BreakInfo ...
func (d Dripstone) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(d)).withBlastResistance(1)
}

// EncodeItem ...
func (d Dripstone) EncodeItem() (name string, meta int16) {
	return "minecraft:dripstone_block", 0
}

// EncodeBlock ...
func (d Dripstone) EncodeBlock() (string, map[string]any) {
	return "minecraft:dripstone_block", nil
}
