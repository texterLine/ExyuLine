// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// Netherrack is a block found in The Nether.
type Netherrack struct {
	solid
	bassDrum
}

// SoilFor ...
func (n Netherrack) SoilFor(block world.Block) bool {
	flower, ok := block.(Flower)
	return ok && flower.Type == WitherRose()
}

// BreakInfo ...
func (n Netherrack) BreakInfo() BreakInfo {
	return newBreakInfo(0.4, pickaxeHarvestable, pickaxeEffective, oneOf(n))
}

// SmeltInfo ...
func (Netherrack) SmeltInfo() item.SmeltInfo {
	return newSmeltInfo(item.NewStack(item.NetherBrick{}, 1), 0.1)
}

// EncodeItem ...
func (Netherrack) EncodeItem() (name string, meta int16) {
	return "minecraft:netherrack", 0
}

// EncodeBlock ...
func (Netherrack) EncodeBlock() (string, map[string]any) {
	return "minecraft:netherrack", nil
}
