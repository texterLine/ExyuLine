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
)

// NetherSprouts are a non-solid plant block that generate in warped forests.
type NetherSprouts struct {
	transparent
	replaceable
	empty
}

// NeighbourUpdateTick ...
func (n NetherSprouts) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !supportsVegetation(n, tx.Block(pos.Side(cube.FaceDown))) {
		breakBlock(n, pos, tx) // TODO: Nylium
	}
}

// UseOnBlock ...
func (n NetherSprouts) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, n)
	if !used {
		return false
	}
	if !supportsVegetation(n, tx.Block(pos.Side(cube.FaceDown))) {
		return false // TODO: Nylium
	}

	place(tx, pos, n, user, ctx)
	return placed(ctx)
}

// HasLiquidDrops ...
func (n NetherSprouts) HasLiquidDrops() bool {
	return false
}

// FlammabilityInfo ...
func (n NetherSprouts) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(0, 0, true)
}

// BreakInfo ...
func (n NetherSprouts) BreakInfo() BreakInfo {
	return newBreakInfo(0, func(t item.Tool) bool {
		return t.ToolType() == item.TypeShears
	}, nothingEffective, oneOf(n))
}

// CompostChance ...
func (NetherSprouts) CompostChance() float64 {
	return 0.5
}

// EncodeItem ...
func (n NetherSprouts) EncodeItem() (name string, meta int16) {
	return "minecraft:nether_sprouts", 0
}

// EncodeBlock ...
func (n NetherSprouts) EncodeBlock() (string, map[string]any) {
	return "minecraft:nether_sprouts", nil
}
