// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// IronBars are blocks that serve a similar purpose to glass panes, but made of iron instead of glass.
type IronBars struct {
	transparent
	thin
	sourceWaterDisplacer

	// Connections holds the sides that the iron bars connects to.
	Connections Connections
}

// NeighbourUpdateTick ...
func (i IronBars) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if connections := calculateConnections(i.Model().(connector), tx, pos); connections != i.Connections {
		i.Connections = connections
		tx.SetBlock(pos, i, nil)
	}
}

// DeriveState ...
func (i IronBars) DeriveState(pos cube.Pos, src world.BlockSource) world.Block {
	i.Connections = calculateConnections(i.Model().(connector), src, pos)
	return i
}

// BreakInfo ...
func (i IronBars) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(i)).withBlastResistance(6)
}

// SideClosed ...
func (i IronBars) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// EncodeItem ...
func (IronBars) EncodeItem() (name string, meta int16) {
	return "minecraft:iron_bars", 0
}

// EncodeBlock ...
func (i IronBars) EncodeBlock() (string, map[string]any) {
	return "minecraft:iron_bars", i.Connections.properties()
}

// allIronBars ...
func allIronBars() (bars []world.Block) {
	for _, c := range allConnections() {
		bars = append(bars, IronBars{Connections: c})
	}
	return
}
