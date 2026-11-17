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

// GlassPane is a transparent block that can be used as a more efficient alternative to glass blocks.
type GlassPane struct {
	transparent
	thin
	clicksAndSticks
	sourceWaterDisplacer

	// Connections holds the sides that the glass pane connects to.
	Connections Connections
}

// NeighbourUpdateTick ...
func (p GlassPane) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if connections := calculateConnections(p.Model().(connector), tx, pos); connections != p.Connections {
		p.Connections = connections
		tx.SetBlock(pos, p, nil)
	}
}

// DeriveState ...
func (p GlassPane) DeriveState(pos cube.Pos, src world.BlockSource) world.Block {
	p.Connections = calculateConnections(p.Model().(connector), src, pos)
	return p
}

// SideClosed ...
func (p GlassPane) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// BreakInfo ...
func (p GlassPane) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, silkTouchOnlyDrop(p))
}

// EncodeItem ...
func (GlassPane) EncodeItem() (name string, meta int16) {
	return "minecraft:glass_pane", meta
}

// EncodeBlock ...
func (p GlassPane) EncodeBlock() (string, map[string]any) {
	return "minecraft:glass_pane", p.Connections.properties()
}

// allGlassPane ...
func allGlassPane() (panes []world.Block) {
	for _, c := range allConnections() {
		panes = append(panes, GlassPane{Connections: c})
	}
	return
}
