// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item/potion"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
	"math"
)

// SplashPotion is an item that grants effects when thrown.
type SplashPotion struct {
	// Type is the type of splash potion.
	Type potion.Potion
}

// MaxCount ...
func (s SplashPotion) MaxCount() int {
	return 1
}

// Use ...
func (s SplashPotion) Use(tx *world.Tx, user User, ctx *UseContext) bool {
	create := tx.World().EntityRegistry().Config().SplashPotion
	opts := world.EntitySpawnOpts{Position: eyePosition(user), Velocity: throwableOffset(user.Rotation()).Vec3().Mul(0.5)}
	tx.AddEntity(create(opts, s.Type, user))
	tx.PlaySound(user.Position(), sound.ItemThrow{})

	ctx.SubtractFromCount(1)
	return true
}

// throwableOffset adds an upwards offset pitch to a throwable entity.
// In vanilla, items such as Splash Potions, Lingering Potions, and
// Bottle o' Enchanting are thrown at a higher angle than where the
// player is looking at.
// The added offset is an ellipse-like shape based on what the input pitch is.
func throwableOffset(r cube.Rotation) cube.Rotation {
	r[1] = max(min(r[1], 89.9), -89.9)
	r[1] -= math.Sqrt(89.9*89.9-r[1]*r[1]) * (26.5 / 89.9)
	r[1] = max(min(r[1], 89.9), -89.9)

	return r
}

// EncodeItem ...
func (s SplashPotion) EncodeItem() (name string, meta int16) {
	return "minecraft:splash_potion", int16(s.Type.Uint8())
}
