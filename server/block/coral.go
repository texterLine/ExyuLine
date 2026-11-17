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
	"math/rand/v2"
	"time"
)

// Coral is a non-solid block that comes in 5 variants.
type Coral struct {
	empty
	transparent
	bassDrum
	sourceWaterDisplacer

	// Type is the type of coral of the block.
	Type CoralType
	// Dead is whether the coral is dead.
	Dead bool
}

// UseOnBlock ...
func (c Coral) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, c)
	if !used {
		return false
	}
	if !tx.Block(pos.Side(cube.FaceDown)).Model().FaceSolid(pos.Side(cube.FaceDown), cube.FaceUp, tx) {
		return false
	}
	if liquid, ok := tx.Liquid(pos); ok {
		if water, ok := liquid.(Water); ok {
			if water.Depth != 8 {
				return false
			}
		}
	}

	place(tx, pos, c, user, ctx)
	return placed(ctx)
}

// HasLiquidDrops ...
func (c Coral) HasLiquidDrops() bool {
	return false
}

// SideClosed ...
func (c Coral) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// NeighbourUpdateTick ...
func (c Coral) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !tx.Block(pos.Side(cube.FaceDown)).Model().FaceSolid(pos.Side(cube.FaceDown), cube.FaceUp, tx) {
		breakBlock(c, pos, tx)
		return
	} else if c.Dead {
		return
	}
	tx.ScheduleBlockUpdate(pos, c, time.Second*5/2)
}

// ScheduledTick ...
func (c Coral) ScheduledTick(pos cube.Pos, tx *world.Tx, _ *rand.Rand) {
	adjacentWater := false
	pos.Neighbours(func(neighbour cube.Pos) {
		if liquid, ok := tx.Liquid(neighbour); ok {
			if _, ok := liquid.(Water); ok {
				adjacentWater = true
			}
		}
	}, tx.Range())
	if !adjacentWater {
		c.Dead = true
		tx.SetBlock(pos, c, nil)
	}
}

// BreakInfo ...
func (c Coral) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, silkTouchOnlyDrop(c))
}

// EncodeBlock ...
func (c Coral) EncodeBlock() (name string, properties map[string]any) {
	if c.Dead {
		return "minecraft:dead_" + c.Type.String() + "_coral", nil
	}
	return "minecraft:" + c.Type.String() + "_coral", nil
}

// EncodeItem ...
func (c Coral) EncodeItem() (name string, meta int16) {
	if c.Dead {
		return "minecraft:dead_" + c.Type.String() + "_coral", 0
	}
	return "minecraft:" + c.Type.String() + "_coral", 0
}

// allCoral returns a list of all coral block variants
func allCoral() (c []world.Block) {
	f := func(dead bool) {
		for _, t := range CoralTypes() {
			c = append(c, Coral{Type: t, Dead: dead})
		}
	}
	f(true)
	f(false)
	return
}
