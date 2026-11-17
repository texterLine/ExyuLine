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

// MuddyMangroveRoots are a decorative variant of mangrove roots.
type MuddyMangroveRoots struct {
	solid

	// Axis is the axis which the muddy mangrove roots faces.
	Axis cube.Axis
}

// BreakInfo ...
func (m MuddyMangroveRoots) BreakInfo() BreakInfo {
	return newBreakInfo(0.7, alwaysHarvestable, shovelEffective, oneOf(m))
}

// SoilFor ...
func (MuddyMangroveRoots) SoilFor(block world.Block) bool {
	switch block.(type) {
	case ShortGrass, Fern, DoubleTallGrass, Flower, DoubleFlower, NetherSprouts, PinkPetals, BambooSapling, Bamboo, Sapling, RedShrub:
		return true
	}
	return false
}

// UseOnBlock ...
func (m MuddyMangroveRoots) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(tx, pos, face, m)
	if !used {
		return
	}
	m.Axis = face.Axis()

	place(tx, pos, m, user, ctx)
	return placed(ctx)
}

// EncodeItem ...
func (MuddyMangroveRoots) EncodeItem() (name string, meta int16) {
	return "minecraft:muddy_mangrove_roots", 0
}

// EncodeBlock ...
func (m MuddyMangroveRoots) EncodeBlock() (string, map[string]any) {
	return "minecraft:muddy_mangrove_roots", map[string]any{"pillar_axis": m.Axis.String()}
}

// allMuddyMangroveRoots ...
func allMuddyMangroveRoots() (roots []world.Block) {
	for _, axis := range cube.Axes() {
		roots = append(roots, MuddyMangroveRoots{Axis: axis})
	}
	return
}
