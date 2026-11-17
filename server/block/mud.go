// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/world"

// Mud is a decorative block obtained by using a water bottle on a dirt block.
type Mud struct {
	solid
}

// SoilFor ...
func (Mud) SoilFor(block world.Block) bool {
	switch block.(type) {
	case ShortGrass, Fern, DoubleTallGrass, Flower, DoubleFlower, NetherSprouts, PinkPetals, DeadBush, BambooSapling, Bamboo, Sapling, RedShrub:
		return true
	}
	return false
}

// BreakInfo ...
func (m Mud) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(m))
}

// EncodeItem ...
func (Mud) EncodeItem() (name string, meta int16) {
	return "minecraft:mud", 0
}

// EncodeBlock ...
func (Mud) EncodeBlock() (string, map[string]any) {
	return "minecraft:mud", nil
}
