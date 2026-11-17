// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// NewFirework creates a firework entity. Firework is an item (and entity) used
// for creating decorative explosions, boosting when flying with elytra, and
// loading into a crossbow as ammunition.
func NewFirework(opts world.EntitySpawnOpts, firework item.Firework) *world.EntityHandle {
	return newFirework(opts, firework, nil, 1.15, 0.04, false)
}

// NewFireworkAttached creates a firework entity with an owner that the firework
// may be attached to.
func NewFireworkAttached(opts world.EntitySpawnOpts, firework item.Firework, owner world.Entity) *world.EntityHandle {
	return newFirework(opts, firework, owner, 0, 0, true)
}

func newFirework(opts world.EntitySpawnOpts, firework item.Firework, owner world.Entity, sidewaysVelocityMultiplier, upwardsAcceleration float64, attached bool) *world.EntityHandle {
	conf := fireworkConf
	conf.SidewaysVelocityMultiplier = sidewaysVelocityMultiplier
	conf.UpwardsAcceleration = upwardsAcceleration
	conf.Firework = firework
	conf.ExistenceDuration = firework.RandomisedDuration()
	conf.Attached = attached
	if attached {
		conf.Owner = owner.H()
	}
	return opts.New(FireworkType, conf)
}

var fireworkConf = FireworkBehaviourConfig{}

// FireworkType is a world.EntityType implementation for Firework.
var FireworkType fireworkType

type fireworkType struct{}

func (t fireworkType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}

func (fireworkType) EncodeEntity() string        { return "minecraft:fireworks_rocket" }
func (fireworkType) BBox(world.Entity) cube.BBox { return cube.BBox{} }

func (fireworkType) DecodeNBT(m map[string]any, data *world.EntityData) {
	conf := fireworkConf
	conf.Firework = item.MapNBT(m, "Item").Item().(item.Firework)
	conf.ExistenceDuration = conf.Firework.RandomisedDuration()

	data.Data = conf.New()
}

func (fireworkType) EncodeNBT(data *world.EntityData) map[string]any {
	return map[string]any{"Item": item.WriteNBT(item.NewStack(data.Data.(*FireworkBehaviour).Firework(), 1), true)}
}
