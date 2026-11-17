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

// SporeBlossom is a decorative block.
type SporeBlossom struct {
	empty
	transparent
}

// HasLiquidDrops ...
func (s SporeBlossom) HasLiquidDrops() bool {
	return true
}

// NeighbourUpdateTick ...
func (s SporeBlossom) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !tx.Block(pos.Side(cube.FaceUp)).Model().FaceSolid(pos.Side(cube.FaceUp), cube.FaceDown, tx) {
		breakBlock(s, pos, tx)
	}
}

// UseOnBlock ...
func (s SporeBlossom) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(tx, pos, face, s)
	if !used {
		return
	}
	if !tx.Block(pos.Side(cube.FaceUp)).Model().FaceSolid(pos.Side(cube.FaceUp), cube.FaceDown, tx) {
		return
	}

	place(tx, pos, s, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (s SporeBlossom) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, oneOf(s))
}

// FlammabilityInfo ...
func (SporeBlossom) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(15, 100, true)
}

// CompostChance ...
func (SporeBlossom) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (s SporeBlossom) EncodeItem() (name string, meta int16) {
	return "minecraft:spore_blossom", 0
}

// EncodeBlock ...
func (s SporeBlossom) EncodeBlock() (string, map[string]any) {
	return "minecraft:spore_blossom", nil
}
