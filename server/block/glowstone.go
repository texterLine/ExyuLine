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
	"github.com/texterLine/exyuline/server/world/sound"
)

// Glowstone is commonly found on the ceiling of the nether dimension.
type Glowstone struct {
	solid
}

func (g Glowstone) Instrument() sound.Instrument {
	return sound.Pling()
}

// CanRedstoneWireStepDown keeps dust from stepping down over glowstone despite its solid top face.
func (Glowstone) CanRedstoneWireStepDown(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

func (Glowstone) RedstoneNonConductive() {}

func (g Glowstone) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, discreteDrops(item.GlowstoneDust{}, g, 2, 4, 4))
}

func (Glowstone) EncodeItem() (name string, meta int16) {
	return "minecraft:glowstone", 0
}

func (Glowstone) EncodeBlock() (string, map[string]any) {
	return "minecraft:glowstone", nil
}

func (Glowstone) LightEmissionLevel() uint8 {
	return 15
}
