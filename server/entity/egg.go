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

// NewEgg creates an Egg entity. Egg is as a throwable entity that can be used
// to spawn chicks.
func NewEgg(opts world.EntitySpawnOpts, owner world.Entity) *world.EntityHandle {
	conf := eggConf
	conf.Owner = owner.H()
	return opts.New(EggType, conf)
}

// TODO: Spawn chicken(e) 12.5% of the time.
var eggConf = ProjectileBehaviourConfig{
	Gravity:       0.03,
	Drag:          0.01,
	Particle:      particle.EggSmash{},
	ParticleCount: 6,
}

// EggType is a world.EntityType implementation for Egg.
var EggType eggType

type eggType struct{}

func (t eggType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}

func (eggType) EncodeEntity() string { return "minecraft:egg" }
func (eggType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

func (eggType) DecodeNBT(_ map[string]any, data *world.EntityData) { data.Data = eggConf.New() }
func (eggType) EncodeNBT(_ *world.EntityData) map[string]any       { return nil }
