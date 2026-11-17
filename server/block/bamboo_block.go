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
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// BambooBlock is a rotatable flammable block made from bamboo.
type BambooBlock struct {
	solid
	bass

	// Axis is the axis which the bamboo block faces.
	Axis cube.Axis
	// Stripped specifies if the bamboo block is stripped.
	Stripped bool
}

// FlammabilityInfo ...
func (BambooBlock) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(5, 5, true)
}

// BreakInfo ...
func (b BambooBlock) BreakInfo() BreakInfo {
	return newBreakInfo(2.0, alwaysHarvestable, axeEffective, oneOf(b))
}

// FuelInfo ...
func (BambooBlock) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// UseOnBlock ...
func (b BambooBlock) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, b)
	if !used {
		return
	}
	b.Axis = face.Axis()

	place(tx, pos, b, user, ctx)
	return placed(ctx)
}

// Strip ...
func (b BambooBlock) Strip() (world.Block, world.Sound, bool) {
	return BambooBlock{Axis: b.Axis, Stripped: true}, nil, !b.Stripped
}

// EncodeItem ...
func (b BambooBlock) EncodeItem() (name string, meta int16) {
	if b.Stripped {
		return "minecraft:stripped_bamboo_block", 0
	}
	return "minecraft:bamboo_block", 0
}

// EncodeBlock ...
func (b BambooBlock) EncodeBlock() (name string, properties map[string]any) {
	meta := map[string]any{"pillar_axis": b.Axis.String()}
	if b.Stripped {
		return "minecraft:stripped_bamboo_block", meta
	}
	return "minecraft:bamboo_block", meta
}

// allBambooBlocks ...
func allBambooBlocks() (blocks []world.Block) {
	for _, axis := range cube.Axes() {
		blocks = append(blocks, BambooBlock{Axis: axis})
		blocks = append(blocks, BambooBlock{Axis: axis, Stripped: true})
	}
	return
}
