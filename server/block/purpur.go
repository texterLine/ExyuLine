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

type (
	// Purpur is a decorative block that is naturally generated in End cities and End ships.
	Purpur struct {
		solid
		bassDrum
	}
	// PurpurPillar is a variant of Purpur that can be rotated.
	PurpurPillar struct {
		solid
		bassDrum

		// Axis is the axis which the purpur pillar block faces.
		Axis cube.Axis
	}
)

// BreakInfo ...
func (p Purpur) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(p)).withBlastResistance(6)
}

// EncodeItem ...
func (p Purpur) EncodeItem() (name string, meta int16) {
	return "minecraft:purpur_block", 0
}

// EncodeBlock ...
func (p Purpur) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:purpur_block", map[string]interface{}{"pillar_axis": "y"}
}

// UseOnBlock ...
func (p PurpurPillar) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, p)
	if !used {
		return
	}
	p.Axis = face.Axis()

	place(tx, pos, p, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (p PurpurPillar) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(p)).withBlastResistance(6)
}

// EncodeItem ...
func (p PurpurPillar) EncodeItem() (name string, meta int16) {
	return "minecraft:purpur_pillar", 0
}

// EncodeBlock ...
func (p PurpurPillar) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:purpur_pillar", map[string]interface{}{"pillar_axis": p.Axis.String()}
}

// allPurpurs ...
func allPurpurs() (purpur []world.Block) {
	purpur = append(purpur, Purpur{})
	for _, axis := range cube.Axes() {
		purpur = append(purpur, PurpurPillar{Axis: axis})
	}
	return
}
