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
	"time"
)

// Loom is a block used to apply patterns on banners. It is also used as a shepherd's job site block that is found in
// villages.
type Loom struct {
	solid
	bass

	// Facing is the direction the loom is facing.
	Facing cube.Direction
}

// FuelInfo ...
func (Loom) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// BreakInfo ...
func (l Loom) BreakInfo() BreakInfo {
	return newBreakInfo(2.5, alwaysHarvestable, axeEffective, oneOf(l))
}

// Activate ...
func (Loom) Activate(pos cube.Pos, _ cube.Face, tx *world.Tx, u item.User, _ *item.UseContext) bool {
	if opener, ok := u.(ContainerOpener); ok {
		opener.OpenBlockContainer(pos, tx)
		return true
	}
	return false
}

// UseOnBlock ...
func (l Loom) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(tx, pos, face, l)
	if !used {
		return
	}
	l.Facing = user.Rotation().Direction().Opposite()
	place(tx, pos, l, user, ctx)
	return placed(ctx)
}

// EncodeItem ...
func (Loom) EncodeItem() (name string, meta int16) {
	return "minecraft:loom", 0
}

// EncodeBlock ...
func (l Loom) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:loom", map[string]interface{}{"direction": int32(horizontalDirection(l.Facing))}
}

// allLooms ...
func allLooms() (looms []world.Block) {
	for _, d := range cube.Directions() {
		looms = append(looms, Loom{Facing: d})
	}
	return
}
