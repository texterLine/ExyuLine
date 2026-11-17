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

// RedstoneBlock is a mineral block equivalent to nine redstone dust.
// It acts as a permanently powered redstone power source that can be pushed by pistons.
type RedstoneBlock struct {
	solid
}

func (r RedstoneBlock) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(r)).withBlastResistance(6)
}

func (r RedstoneBlock) EncodeItem() (name string, meta int16) {
	return "minecraft:redstone_block", 0
}

func (r RedstoneBlock) EncodeBlock() (string, map[string]any) {
	return "minecraft:redstone_block", nil
}

func (r RedstoneBlock) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, r)
	if !used {
		return false
	}
	place(tx, pos, r, user, ctx)
	return placed(ctx)
}

func (RedstoneBlock) RedstonePower(cube.Pos, *world.Tx, cube.Face) int {
	return 15
}

// RedstoneStrongPower returns no strong power. Redstone blocks power adjacent components directly, but do not strongly
// power adjacent opaque blocks.
func (RedstoneBlock) RedstoneStrongPower(cube.Pos, *world.Tx, cube.Face) int {
	return 0
}

func (RedstoneBlock) RedstoneNonConductive() {}
