// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// PolishedTuff is a decorational variant of Tuff that can be crafted or found naturally in Trial Chambers.
type PolishedTuff struct {
	solid
	bassDrum
}

// BreakInfo ...
func (t PolishedTuff) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(t)).withBlastResistance(6)
}

// EncodeItem ...
func (t PolishedTuff) EncodeItem() (name string, meta int16) {
	return "minecraft:polished_tuff", 0
}

// EncodeBlock ...
func (t PolishedTuff) EncodeBlock() (string, map[string]any) {
	return "minecraft:polished_tuff", nil
}
