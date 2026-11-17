// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/item/potion"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
)

// LingeringPotion is a variant of a splash potion that can be thrown to leave clouds with status effects that linger on
// the ground in an area.
type LingeringPotion struct {
	// Type is the type of lingering potion.
	Type potion.Potion
}

// MaxCount ...
func (l LingeringPotion) MaxCount() int {
	return 1
}

// Use ...
func (l LingeringPotion) Use(tx *world.Tx, user User, ctx *UseContext) bool {
	create := tx.World().EntityRegistry().Config().LingeringPotion
	opts := world.EntitySpawnOpts{Position: eyePosition(user), Velocity: throwableOffset(user.Rotation()).Vec3().Mul(0.5)}
	tx.AddEntity(create(opts, l.Type, user))
	tx.PlaySound(user.Position(), sound.ItemThrow{})

	ctx.SubtractFromCount(1)
	return true
}

// EncodeItem ...
func (l LingeringPotion) EncodeItem() (name string, meta int16) {
	return "minecraft:lingering_potion", int16(l.Type.Uint8())
}
