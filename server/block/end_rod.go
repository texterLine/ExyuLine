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

// EndRod is a decorative light source that emits white particles.
type EndRod struct {
	transparent
	flowingWaterDisplacer

	// Facing is the direction the white end point of the end rod is facing.
	Facing cube.Face
}

// UseOnBlock ...
func (e EndRod) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, face, used := firstReplaceable(tx, pos, face, e)
	if !used {
		return false
	}

	e.Facing = face
	if other, ok := tx.Block(pos.Side(face.Opposite())).(EndRod); ok {
		if face == other.Facing {
			e.Facing = face.Opposite()
		}
	}
	place(tx, pos, e, user, ctx)
	return placed(ctx)
}

// SideClosed ...
func (EndRod) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// Model ...
func (e EndRod) Model() world.BlockModel {
	return model.EndRod{Axis: e.Facing.Axis()}
}

// LightEmissionLevel ...
func (EndRod) LightEmissionLevel() uint8 {
	return 14
}

// BreakInfo ...
func (e EndRod) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, oneOf(e))
}

// EncodeItem ...
func (EndRod) EncodeItem() (name string, meta int16) {
	return "minecraft:end_rod", 0
}

// EncodeBlock ...
func (e EndRod) EncodeBlock() (string, map[string]any) {
	if e.Facing.Axis() == cube.Y {
		return "minecraft:end_rod", map[string]any{"facing_direction": int32(e.Facing)}
	}
	return "minecraft:end_rod", map[string]any{"facing_direction": int32(e.Facing.Opposite())}
}

// allEndRods ...
func allEndRods() (b []world.Block) {
	for _, f := range cube.Faces() {
		b = append(b, EndRod{Facing: f})
	}
	return
}
