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
	"math/rand/v2"
)

// DeadBush is a transparent block in the form of an aesthetic plant.
type DeadBush struct {
	empty
	replaceable
	transparent
	sourceWaterDisplacer
}

// NeighbourUpdateTick ...
func (d DeadBush) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !supportsVegetation(d, tx.Block(pos.Side(cube.FaceDown))) {
		breakBlock(d, pos, tx)
	}
}

// UseOnBlock ...
func (d DeadBush) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, d)
	if !used || !supportsVegetation(d, tx.Block(pos.Side(cube.FaceDown))) {
		return false
	}

	place(tx, pos, d, user, ctx)
	return placed(ctx)
}

// SideClosed ...
func (d DeadBush) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// HasLiquidDrops ...
func (d DeadBush) HasLiquidDrops() bool {
	return true
}

// FlammabilityInfo ...
func (d DeadBush) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(60, 100, true)
}

// BreakInfo ...
func (d DeadBush) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
		if t.ToolType() == item.TypeShears {
			return []item.Stack{item.NewStack(d, 1)}
		}
		if amount := rand.IntN(3); amount != 0 {
			return []item.Stack{item.NewStack(item.Stick{}, amount)}
		}
		return nil
	})
}

// EncodeItem ...
func (d DeadBush) EncodeItem() (name string, meta int16) {
	return "minecraft:deadbush", 0
}

// EncodeBlock ...
func (d DeadBush) EncodeBlock() (string, map[string]any) {
	return "minecraft:deadbush", nil
}
