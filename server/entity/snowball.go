// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/particle"
)

// NewSnowball creates a snowball entity at a position with an owner entity.
func NewSnowball(opts world.EntitySpawnOpts, owner world.Entity) *world.EntityHandle {
	conf := snowballConf
	conf.Owner = owner.H()
	return opts.New(SnowballType, conf)
}

var snowballConf = ProjectileBehaviourConfig{
	Gravity:       0.03,
	Drag:          0.01,
	Particle:      particle.SnowballPoof{},
	ParticleCount: 6,
}

// SnowballType is a world.EntityType implementation for snowballs.
var SnowballType snowballType

type snowballType struct{}

func (t snowballType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}

func (snowballType) EncodeEntity() string { return "minecraft:snowball" }
func (snowballType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

func (snowballType) DecodeNBT(_ map[string]any, data *world.EntityData) {
	data.Data = snowballConf.New()
}
func (snowballType) EncodeNBT(*world.EntityData) map[string]any { return nil }
