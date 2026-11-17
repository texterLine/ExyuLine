// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// QuartzBricks is a mineral block used only for decoration.
type QuartzBricks struct {
	solid
	bassDrum
}

// BreakInfo ...
func (q QuartzBricks) BreakInfo() BreakInfo {
	return newBreakInfo(0.8, pickaxeHarvestable, pickaxeEffective, oneOf(q))
}

// EncodeItem ...
func (QuartzBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:quartz_bricks", 0
}

// EncodeBlock ...
func (QuartzBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:quartz_bricks", nil
}
