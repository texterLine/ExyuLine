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

// Pumpkin is a crop block. Interacting with shears results in the carved variant.
type Pumpkin struct {
	solid

	// Carved is whether the pumpkin is carved.
	Carved bool
	// Facing is the direction the pumpkin is facing.
	Facing cube.Direction
}

// Instrument ...
func (p Pumpkin) Instrument() sound.Instrument {
	if !p.Carved {
		return sound.Didgeridoo()
	}
	return sound.Piano()
}

// UseOnBlock ...
func (p Pumpkin) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(tx, pos, face, p)
	if !used {
		return
	}
	p.Facing = user.Rotation().Direction().Opposite()

	place(tx, pos, p, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (p Pumpkin) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, axeEffective, oneOf(p))
}

// CompostChance ...
func (Pumpkin) CompostChance() float64 {
	return 0.65
}

// Carve ...
func (p Pumpkin) Carve(f cube.Face) (world.Block, bool) {
	return Pumpkin{Facing: f.Direction(), Carved: true}, !p.Carved
}

// Helmet ...
func (p Pumpkin) Helmet() bool {
	return p.Carved
}

// DefencePoints ...
func (p Pumpkin) DefencePoints() float64 {
	return 0
}

// Toughness ...
func (p Pumpkin) Toughness() float64 {
	return 0
}

// KnockBackResistance ...
func (p Pumpkin) KnockBackResistance() float64 {
	return 0
}

// EncodeItem ...
func (p Pumpkin) EncodeItem() (name string, meta int16) {
	if p.Carved {
		return "minecraft:carved_pumpkin", 0
	}
	return "minecraft:pumpkin", 0
}

// EncodeBlock ...
func (p Pumpkin) EncodeBlock() (name string, properties map[string]any) {
	if p.Carved {
		return "minecraft:carved_pumpkin", map[string]any{"minecraft:cardinal_direction": p.Facing.String()}
	}
	return "minecraft:pumpkin", map[string]any{"minecraft:cardinal_direction": p.Facing.String()}
}

func allPumpkins() (pumpkins []world.Block) {
	for i := cube.Direction(0); i <= 3; i++ {
		pumpkins = append(pumpkins, Pumpkin{Facing: i})
		pumpkins = append(pumpkins, Pumpkin{Facing: i, Carved: true})
	}
	return
}
