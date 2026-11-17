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

// StainedGlassPane is a transparent block that can be used as a more efficient alternative to glass blocks.
type StainedGlassPane struct {
	transparent
	thin
	clicksAndSticks
	sourceWaterDisplacer

	// Colour specifies the colour of the block.
	Colour item.Colour
	// Connections holds the sides that the glass pane connects to.
	Connections Connections
}

// NeighbourUpdateTick ...
func (p StainedGlassPane) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if connections := calculateConnections(p.Model().(connector), tx, pos); connections != p.Connections {
		p.Connections = connections
		tx.SetBlock(pos, p, nil)
	}
}

// DeriveState ...
func (p StainedGlassPane) DeriveState(pos cube.Pos, src world.BlockSource) world.Block {
	p.Connections = calculateConnections(p.Model().(connector), src, pos)
	return p
}

// SideClosed ...
func (p StainedGlassPane) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// BreakInfo ...
func (p StainedGlassPane) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, silkTouchOnlyDrop(p))
}

// EncodeItem ...
func (p StainedGlassPane) EncodeItem() (name string, meta int16) {
	return "minecraft:" + p.Colour.String() + "_stained_glass_pane", 0
}

// EncodeBlock ...
func (p StainedGlassPane) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + p.Colour.String() + "_stained_glass_pane", p.Connections.properties()
}

// allStainedGlassPane returns stained-glass panes with all possible colours.
func allStainedGlassPane() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, colour := range item.Colours() {
		for _, c := range allConnections() {
			b = append(b, StainedGlassPane{Colour: colour, Connections: c})
		}
	}
	return b
}
