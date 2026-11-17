// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// Slime is a storage block equivalent to nine slimeballs. It has both sticky and bouncy properties,
// making it useful in conjunction with pistons to move both blocks and entities.
type Slime struct {
	solid
	transparent
}

// EntityLand ...
func (Slime) EntityLand(_ cube.Pos, _ *world.Tx, e world.Entity, distance *float64) {
	if _, ok := e.(fallDistanceEntity); ok {
		*distance = 0
	}
	if s, ok := e.(interface{ Sneaking() bool }); ok && s.Sneaking() {
		return
	}
	if v, ok := e.(velocityEntity); ok {
		vel := v.Velocity()
		if vel[1] < 0 {
			vel[1] = -vel[1]
			v.SetVelocity(vel)
		}
	}
}

// Friction ...
func (Slime) Friction() float64 {
	return 0.8
}

// BreakInfo ...
func (s Slime) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, oneOf(s))
}

// EncodeItem ...
func (Slime) EncodeItem() (name string, meta int16) {
	return "minecraft:slime", 0
}

// EncodeBlock ...
func (Slime) EncodeBlock() (string, map[string]any) {
	return "minecraft:slime", nil
}
