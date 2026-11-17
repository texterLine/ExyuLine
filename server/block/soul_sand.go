// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
)

// SoulSand is a block found naturally only in the Nether. SoulSand slows movement of mobs & players.
type SoulSand struct {
	solid
}

// TODO: Implement bubble columns.

// SoilFor ...
func (s SoulSand) SoilFor(block world.Block) bool {
	flower, ok := block.(Flower)
	return ok && flower.Type == WitherRose()
}

// Instrument ...
func (s SoulSand) Instrument() sound.Instrument {
	return sound.CowBell()
}

// BreakInfo ...
func (s SoulSand) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(s))
}

// EncodeItem ...
func (SoulSand) EncodeItem() (name string, meta int16) {
	return "minecraft:soul_sand", 0
}

// EncodeBlock ...
func (SoulSand) EncodeBlock() (string, map[string]any) {
	return "minecraft:soul_sand", nil
}
