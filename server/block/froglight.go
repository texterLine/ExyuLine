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

// Froglight is a luminous natural block that can be obtained if a frog eats a tiny magma cube.
type Froglight struct {
	solid

	// Type is the type of froglight.
	Type FroglightType
	// Axis is the axis which the froglight block faces.
	Axis cube.Axis
}

// LightEmissionLevel ...
func (f Froglight) LightEmissionLevel() uint8 {
	return 15
}

// UseOnBlock ...
func (f Froglight) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, f)
	if !used {
		return
	}
	f.Axis = face.Axis()

	place(tx, pos, f, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (f Froglight) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, oneOf(f))
}

// EncodeItem ...
func (f Froglight) EncodeItem() (name string, meta int16) {
	return "minecraft:" + f.Type.String() + "_froglight", 0
}

// EncodeBlock ...
func (f Froglight) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + f.Type.String() + "_froglight", map[string]any{"pillar_axis": f.Axis.String()}
}

// allFrogLight ...
func allFroglight() (froglight []world.Block) {
	for _, axis := range cube.Axes() {
		for _, t := range FroglightTypes() {
			froglight = append(froglight, Froglight{Type: t, Axis: axis})
		}
	}
	return
}
