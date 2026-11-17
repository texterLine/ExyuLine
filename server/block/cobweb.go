// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// Cobweb is a non-solid block that drastically slows entities passing through it. It is broken
// quickly with a sword or shears and drops string when broken without silk touch.
type Cobweb struct {
	empty
	transparent
}

// Cobweb is implemented because the item package needs to identify this block but cannot implement the block package.
func (Cobweb) Cobweb() {}

// EntityInside slows the entity's velocity and resets its fall distance while it is inside the cobweb.
func (Cobweb) EntityInside(_ cube.Pos, _ *world.Tx, e world.Entity) {
	if fallEntity, ok := e.(fallDistanceEntity); ok {
		fallEntity.ResetFallDistance()
	}
	if v, ok := e.(velocityEntity); ok {
		vel := v.Velocity()
		vel[0] *= 0.25
		vel[1] *= 0.05
		vel[2] *= 0.25
		v.SetVelocity(vel)
	}
}

// BreakInfo ...
func (c Cobweb) BreakInfo() BreakInfo {
	return newBreakInfo(
		4,
		alwaysHarvestable,
		func(t item.Tool) bool {
			return swordEffective(t) || shearsEffective(t)
		},
		func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
			if t.ToolType() == item.TypeShears {
				return oneOf(c)(t, enchantments)
			}
			if t.ToolType() == item.TypeSword {
				return oneOf(String{})(t, enchantments)
			}
			return nil
		},
	)
}

// HasLiquidDrops ...
func (Cobweb) HasLiquidDrops() bool {
	return true
}

// EncodeItem ...
func (Cobweb) EncodeItem() (name string, meta int16) {
	return "minecraft:web", 0
}

// EncodeBlock ...
func (Cobweb) EncodeBlock() (string, map[string]any) {
	return "minecraft:web", nil
}
