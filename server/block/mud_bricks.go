// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// MudBricks are a decorative block obtained by crafting 4 packed mud.
type MudBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (m MudBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, alwaysHarvestable, nothingEffective, oneOf(m)).withBlastResistance(3)
}

// EncodeItem ...
func (MudBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:mud_bricks", 0
}

// EncodeBlock ...
func (MudBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:mud_bricks", nil
}
