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

// InfestedDeepslate is a block that hides a silverfish. It looks identical to deepslate.
// TODO: spawn a silverfish on break (without silk touch) once silverfish are implemented.
type InfestedDeepslate struct {
	solid
	flute

	// Axis is the axis which the deepslate faces.
	Axis cube.Axis
}

// BreakInfo ...
func (i InfestedDeepslate) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, silkTouchOnlyDrop(i)).withBlastResistance(0.75)
}

// UseOnBlock ...
func (i InfestedDeepslate) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, i)
	if !used {
		return
	}
	i.Axis = face.Axis()

	place(tx, pos, i, user, ctx)
	return placed(ctx)
}

// EncodeItem ...
func (InfestedDeepslate) EncodeItem() (name string, meta int16) {
	return "minecraft:infested_deepslate", 0
}

// EncodeBlock ...
func (i InfestedDeepslate) EncodeBlock() (string, map[string]any) {
	return "minecraft:infested_deepslate", map[string]any{"pillar_axis": i.Axis.String()}
}

// allInfestedDeepslate ...
func allInfestedDeepslate() (s []world.Block) {
	for _, axis := range cube.Axes() {
		s = append(s, InfestedDeepslate{Axis: axis})
	}
	return
}
