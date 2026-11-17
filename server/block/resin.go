// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Resin is a block equivalent to nine resin clumps.
type Resin struct {
	solid
}

// BreakInfo ...
func (r Resin) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, oneOf(r))
}

// EncodeItem ...
func (Resin) EncodeItem() (name string, meta int16) {
	return "minecraft:resin_block", 0
}

// EncodeBlock ...
func (Resin) EncodeBlock() (string, map[string]any) {
	return "minecraft:resin_block", nil
}
