// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/world"
)

// InfestedStoneBricks is a block that hides a silverfish. It looks identical to stone bricks.
// TODO: spawn a silverfish on break (without silk touch) once silverfish are implemented.
type InfestedStoneBricks struct {
	solid
	flute

	// Type is the type of stone bricks of the block.
	Type StoneBricksType
}

// BreakInfo ...
func (i InfestedStoneBricks) BreakInfo() BreakInfo {
	return newBreakInfo(0.75, pickaxeHarvestable, pickaxeEffective, silkTouchOnlyDrop(i)).withBlastResistance(0.75)
}

// EncodeItem ...
func (i InfestedStoneBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:infested_" + i.Type.String(), 0
}

// EncodeBlock ...
func (i InfestedStoneBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:infested_" + i.Type.String(), nil
}

// allInfestedStoneBricks returns a list of all infested stone bricks variants.
func allInfestedStoneBricks() (s []world.Block) {
	for _, t := range StoneBricksTypes() {
		s = append(s, InfestedStoneBricks{Type: t})
	}
	return
}
