// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/world"

// SoulSoil is a block naturally found only in the soul sand valley.
type SoulSoil struct {
	solid
}

// SoilFor ...
func (s SoulSoil) SoilFor(block world.Block) bool {
	_, ok := block.(NetherSprouts)
	return ok
}

// BreakInfo ...
func (s SoulSoil) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(s))
}

// EncodeItem ...
func (SoulSoil) EncodeItem() (name string, meta int16) {
	return "minecraft:soul_soil", 0
}

// EncodeBlock ...
func (SoulSoil) EncodeBlock() (string, map[string]any) {
	return "minecraft:soul_soil", nil
}
