// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Amethyst is a decorative block crafted from four amethyst shards.
type Amethyst struct {
	solid
}

// BreakInfo ...
func (a Amethyst) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeHarvestable, oneOf(a))
}

// EncodeItem ...
func (Amethyst) EncodeItem() (name string, meta int16) {
	return "minecraft:amethyst_block", 0
}

// EncodeBlock ...
func (Amethyst) EncodeBlock() (string, map[string]any) {
	return "minecraft:amethyst_block", nil
}
