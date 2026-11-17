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

// NetherBricks are blocks used to form nether fortresses in the Nether.
// Red Nether bricks, Cracked Nether bricks and Chiseled Nether bricks are decorative variants that do not naturally generate.
type NetherBricks struct {
	solid
	bassDrum

	// NetherBricksType is the type of nether bricks of the block.
	Type NetherBricksType
}

// BreakInfo ...
func (n NetherBricks) BreakInfo() BreakInfo {
	return newBreakInfo(2, pickaxeHarvestable, pickaxeEffective, oneOf(n)).withBlastResistance(6)
}

// SmeltInfo ...
func (n NetherBricks) SmeltInfo() item.SmeltInfo {
	if n.Type == NormalNetherBricks() {
		return newSmeltInfo(item.NewStack(NetherBricks{Type: CrackedNetherBricks()}, 1), 0.1)
	}
	return item.SmeltInfo{}
}

// EncodeItem ...
func (n NetherBricks) EncodeItem() (id string, meta int16) {
	return "minecraft:" + n.Type.String(), 0
}

// EncodeBlock ...
func (n NetherBricks) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + n.Type.String(), nil
}

// allNetherBricks returns a list of all nether bricks variants.
func allNetherBricks() (netherBricks []world.Block) {
	for _, t := range NetherBricksTypes() {
		netherBricks = append(netherBricks, NetherBricks{Type: t})
	}
	return
}
