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

// Fern is a transparent plant block which can be used to obtain seeds and as decoration.
type Fern struct {
	replaceable
	transparent
	empty
}

// FlammabilityInfo ...
func (g Fern) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(60, 100, false)
}

// BreakInfo ...
func (g Fern) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, grassDrops(g))
}

// BoneMeal attempts to affect the block using a bone meal item.
func (g Fern) BoneMeal(pos cube.Pos, tx *world.Tx) item.BoneMealResult {
	upper := DoubleTallGrass{Type: FernDoubleTallGrass(), UpperPart: true}
	if replaceableWith(tx, pos.Side(cube.FaceUp), upper) {
		tx.SetBlock(pos, DoubleTallGrass{Type: FernDoubleTallGrass()}, nil)
		tx.SetBlock(pos.Side(cube.FaceUp), upper, nil)
		return item.BoneMealResultSmall
	}
	return item.BoneMealResultNone
}

// CompostChance ...
func (g Fern) CompostChance() float64 {
	return 0.3
}

// NeighbourUpdateTick ...
func (g Fern) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !supportsVegetation(g, tx.Block(pos.Side(cube.FaceDown))) {
		breakBlock(g, pos, tx)
	}
}

// HasLiquidDrops ...
func (g Fern) HasLiquidDrops() bool {
	return true
}

// UseOnBlock ...
func (g Fern) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, g)
	if !used || !supportsVegetation(g, tx.Block(pos.Side(cube.FaceDown))) {
		return false
	}

	place(tx, pos, g, user, ctx)
	return placed(ctx)
}

// EncodeItem ...
func (g Fern) EncodeItem() (name string, meta int16) {
	return "minecraft:fern", 0
}

// EncodeBlock ...
func (g Fern) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:fern", nil
}
