// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// InfestedCobblestone is a block that hides a silverfish. It looks identical to cobblestone.
// TODO: spawn a silverfish on break (without silk touch) once silverfish are implemented.
type InfestedCobblestone struct {
	solid
	flute
}

// BreakInfo ...
func (i InfestedCobblestone) BreakInfo() BreakInfo {
	return newBreakInfo(1, pickaxeHarvestable, pickaxeEffective, silkTouchOnlyDrop(i)).withBlastResistance(0.75)
}

// EncodeItem ...
func (InfestedCobblestone) EncodeItem() (name string, meta int16) {
	return "minecraft:infested_cobblestone", 0
}

// EncodeBlock ...
func (InfestedCobblestone) EncodeBlock() (string, map[string]any) {
	return "minecraft:infested_cobblestone", nil
}
