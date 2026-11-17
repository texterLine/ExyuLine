// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/world"

// Podzol is a dirt-type block that naturally blankets the surface of the giant tree taiga and bamboo jungles, along
// with their respective variants.
type Podzol struct {
	solid
}

// SoilFor ...
func (p Podzol) SoilFor(block world.Block) bool {
	switch block.(type) {
	case ShortGrass, Fern, DoubleTallGrass, Flower, DoubleFlower, NetherSprouts, DeadBush, SugarCane, BambooSapling, Bamboo, Sapling, RedShrub:
		return true
	}
	return false
}

// Shovel ...
func (Podzol) Shovel() (world.Block, bool) {
	return DirtPath{}, true
}

// BreakInfo ...
func (p Podzol) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, silkTouchOneOf(Dirt{}, p))
}

// EncodeItem ...
func (Podzol) EncodeItem() (name string, meta int16) {
	return "minecraft:podzol", 0
}

// EncodeBlock ...
func (Podzol) EncodeBlock() (string, map[string]any) {
	return "minecraft:podzol", nil
}
