// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/block/model"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// IronChain is a metallic decoration block.
type IronChain struct {
	transparent
	sourceWaterDisplacer

	// Axis is the axis which the chain faces.
	Axis cube.Axis
}

// SideClosed ...
func (IronChain) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// UseOnBlock ...
func (c IronChain) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, c)
	if !used {
		return
	}
	c.Axis = face.Axis()

	place(tx, pos, c, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (c IronChain) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// EncodeItem ...
func (IronChain) EncodeItem() (name string, meta int16) {
	return "minecraft:iron_chain", 0
}

// EncodeBlock ...
func (c IronChain) EncodeBlock() (string, map[string]any) {
	return "minecraft:iron_chain", map[string]any{"pillar_axis": c.Axis.String()}
}

// Model ...
func (c IronChain) Model() world.BlockModel {
	return model.Chain{Axis: c.Axis}
}

// allIronChains ...
func allIronChains() (chains []world.Block) {
	for _, axis := range cube.Axes() {
		chains = append(chains, IronChain{Axis: axis})
	}
	return
}
