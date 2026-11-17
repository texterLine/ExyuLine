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

// LitPumpkin is a decorative light emitting block crafted with a Carved Pumpkin & Torch
type LitPumpkin struct {
	solid

	// Facing is the direction the pumpkin is facing.
	Facing cube.Direction
}

// LightEmissionLevel ...
func (l LitPumpkin) LightEmissionLevel() uint8 {
	return 15
}

// UseOnBlock ...
func (l LitPumpkin) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(tx, pos, face, l)
	if !used {
		return
	}
	l.Facing = user.Rotation().Direction().Opposite()

	place(tx, pos, l, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (l LitPumpkin) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, axeEffective, oneOf(l))
}

// EncodeItem ...
func (l LitPumpkin) EncodeItem() (name string, meta int16) {
	return "minecraft:lit_pumpkin", 0
}

// EncodeBlock ...
func (l LitPumpkin) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:lit_pumpkin", map[string]any{"minecraft:cardinal_direction": l.Facing.String()}
}

func allLitPumpkins() (pumpkins []world.Block) {
	for i := cube.Direction(0); i <= 3; i++ {
		pumpkins = append(pumpkins, LitPumpkin{Facing: i})
	}
	return
}
