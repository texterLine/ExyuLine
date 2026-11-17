// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/internal/nbtconv"
	"github.com/texterLine/exyuline/server/world"
)

// NewFallingBlock creates a new FallingBlock entity.
func NewFallingBlock(opts world.EntitySpawnOpts, block world.Block) *world.EntityHandle {
	conf := fallingBlockConf
	conf.Block = block
	return opts.New(FallingBlockType, conf)
}

var fallingBlockConf = FallingBlockBehaviourConfig{
	Gravity: 0.04,
	Drag:    0.02,
}

// FallingBlockType is a world.EntityType implementation for FallingBlock.
var FallingBlockType fallingBlockType

type fallingBlockType struct{}

func (t fallingBlockType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}
func (fallingBlockType) EncodeEntity() string   { return "minecraft:falling_block" }
func (fallingBlockType) NetworkOffset() float64 { return 0.49 }
func (fallingBlockType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 0.98, 0.49)
}

func (fallingBlockType) DecodeNBT(m map[string]any, data *world.EntityData) {
	conf := fallingBlockConf
	conf.Block = nbtconv.Block(m, "FallingBlock")
	conf.DistanceFallen = nbtconv.Float64(m, "FallDistance")
	data.Data = conf.New()
}

func (fallingBlockType) EncodeNBT(data *world.EntityData) map[string]any {
	b := data.Data.(*FallingBlockBehaviour)
	return map[string]any{"FallDistance": b.passive.fallDistance, "FallingBlock": nbtconv.WriteBlock(b.block)}
}
