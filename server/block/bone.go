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
	"github.com/texterLine/exyuline/server/world/sound"
)

// Bone is a decorative block that can face different directions.
type Bone struct {
	solid

	// Axis is the axis which the bone block faces.
	Axis cube.Axis
}

// Instrument ...
func (b Bone) Instrument() sound.Instrument {
	return sound.Xylophone()
}

// UseOnBlock handles the rotational placing of bone blocks.
func (b Bone) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, b)
	if !used {
		return
	}
	b.Axis = face.Axis()

	place(tx, pos, b, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (b Bone) BreakInfo() BreakInfo {
	return newBreakInfo(2, pickaxeHarvestable, pickaxeEffective, oneOf(b))
}

// EncodeItem ...
func (b Bone) EncodeItem() (name string, meta int16) {
	return "minecraft:bone_block", 0
}

// EncodeBlock ...
func (b Bone) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:bone_block", map[string]any{"pillar_axis": b.Axis.String(), "deprecated": int32(0)}
}

// allBoneBlock ...
func allBoneBlock() (boneBlocks []world.Block) {
	for _, axis := range cube.Axes() {
		boneBlocks = append(boneBlocks, Bone{Axis: axis})
	}
	return
}
