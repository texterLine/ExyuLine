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

// LilyPad is a short, flat solid block that can be found naturally growing only on water, in swamps and wheat
// farm rooms in woodland mansions.
type LilyPad struct {
	transparent
}

// HasLiquidDrops ...
func (LilyPad) HasLiquidDrops() bool {
	return true
}

// CompostChance ...
func (LilyPad) CompostChance() float64 {
	return 0.65
}

// NeighbourUpdateTick ...
func (l LilyPad) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if liq, ok := tx.Liquid(pos.Side(cube.FaceDown)); !ok || liq.LiquidType() != "water" || liq.LiquidDepth() < 8 {
		breakBlock(l, pos, tx)
	}
}

// UseOnBlock ...
func (l LilyPad) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, l)
	if !used {
		return false
	}
	if liq, ok := tx.Liquid(pos.Side(cube.FaceDown)); !ok || liq.LiquidType() != "water" || liq.LiquidDepth() < 8 {
		return false
	}
	place(tx, pos, l, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (l LilyPad) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, oneOf(l))
}

// EncodeItem ...
func (LilyPad) EncodeItem() (name string, meta int16) {
	return "minecraft:waterlily", 0
}

// Model ...
func (LilyPad) Model() world.BlockModel {
	return model.LilyPad{}
}

// EncodeBlock ...
func (LilyPad) EncodeBlock() (string, map[string]any) {
	return "minecraft:waterlily", nil
}
