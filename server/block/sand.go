// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// Sand is a block affected by gravity. It can come in a red variant.
type Sand struct {
	gravityAffected
	solid
	snare

	// Red toggles the red sand variant.
	Red bool
}

// SoilFor ...
func (s Sand) SoilFor(block world.Block) bool {
	switch block.(type) {
	case Cactus, DeadBush, SugarCane, BambooSapling, Bamboo:
		return true
	}
	return false
}

// NeighbourUpdateTick ...
func (s Sand) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	s.fall(s, pos, tx)
}

// BreakInfo ...
func (s Sand) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(s))
}

// SmeltInfo ...
func (Sand) SmeltInfo() item.SmeltInfo {
	return newSmeltInfo(item.NewStack(Glass{}, 1), 0.1)
}

// EncodeItem ...
func (s Sand) EncodeItem() (name string, meta int16) {
	if s.Red {
		return "minecraft:red_sand", 0
	}
	return "minecraft:sand", 0
}

// EncodeBlock ...
func (s Sand) EncodeBlock() (string, map[string]any) {
	if s.Red {
		return "minecraft:red_sand", nil
	}
	return "minecraft:sand", nil
}
