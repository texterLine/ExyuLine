// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
	"time"
)

// EnderPearl is a smooth, greenish-blue item used to teleport and to make an eye of ender.
type EnderPearl struct{}

// Use ...
func (e EnderPearl) Use(tx *world.Tx, user User, ctx *UseContext) bool {
	create := tx.World().EntityRegistry().Config().EnderPearl
	opts := world.EntitySpawnOpts{Position: eyePosition(user), Velocity: user.Rotation().Vec3().Mul(1.5)}
	tx.AddEntity(create(opts, user))
	tx.PlaySound(user.Position(), sound.ItemThrow{})

	ctx.SubtractFromCount(1)
	return true
}

// Cooldown ...
func (EnderPearl) Cooldown() time.Duration {
	return time.Second
}

// MaxCount ...
func (EnderPearl) MaxCount() int {
	return 16
}

// EncodeItem ...
func (EnderPearl) EncodeItem() (name string, meta int16) {
	return "minecraft:ender_pearl", 0
}
