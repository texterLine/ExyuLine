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

// HayBale is a decorative, flammable block that can also be used to
// feed horses, breed llamas, reduce fall damage, and extend campfire smokes.
type HayBale struct {
	solid

	// Axis is the axis which the hay bale block faces.
	Axis cube.Axis
}

// Instrument ...
func (HayBale) Instrument() sound.Instrument {
	return sound.Banjo()
}

// EntityLand ...
func (h HayBale) EntityLand(_ cube.Pos, _ *world.Tx, e world.Entity, distance *float64) {
	if _, ok := e.(fallDistanceEntity); ok {
		*distance *= 0.2
	}
}

// UseOnBlock ...
func (h HayBale) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, h)
	if !used {
		return
	}
	h.Axis = face.Axis()

	place(tx, pos, h, user, ctx)
	return placed(ctx)
}

// FlammabilityInfo ...
func (HayBale) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(60, 20, false)
}

// BreakInfo ...
func (h HayBale) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, hoeEffective, oneOf(h))
}

// CompostChance ...
func (HayBale) CompostChance() float64 {
	return 0.85
}

// EncodeItem ...
func (HayBale) EncodeItem() (name string, meta int16) {
	return "minecraft:hay_block", 0
}

// EncodeBlock ...
func (h HayBale) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:hay_block", map[string]interface{}{"pillar_axis": h.Axis.String(), "deprecated": int32(0)}
}

// allHayBales ...
func allHayBales() (haybale []world.Block) {
	for _, a := range cube.Axes() {
		haybale = append(haybale, HayBale{Axis: a})
	}
	return
}
