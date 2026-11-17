// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"math/rand/v2"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// NetherWart is a fungus found in the Nether that is vital in the creation of potions.
type NetherWart struct {
	transparent
	empty

	// Age is the age of the nether wart block. 3 is fully grown.
	Age int
}

// HasLiquidDrops ...
func (n NetherWart) HasLiquidDrops() bool {
	return true
}

// RandomTick ...
func (n NetherWart) RandomTick(pos cube.Pos, tx *world.Tx, r *rand.Rand) {
	if n.Age < 3 && r.Float64() < 0.1 {
		n.Age++
		tx.SetBlock(pos, n, nil)
	}
}

// UseOnBlock ...
func (n NetherWart) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, n)
	if !used {
		return false
	}
	if _, ok := tx.Block(pos.Side(cube.FaceDown)).(SoulSand); !ok {
		return false
	}

	place(tx, pos, n, user, ctx)
	return placed(ctx)
}

// NeighbourUpdateTick ...
func (n NetherWart) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if _, ok := tx.Block(pos.Side(cube.FaceDown)).(SoulSand); !ok {
		breakBlock(n, pos, tx)
	}
}

// BreakInfo ...
func (n NetherWart) BreakInfo() BreakInfo {
	return newBreakInfo(0, alwaysHarvestable, nothingEffective, func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
		if n.Age < 3 {
			return []item.Stack{item.NewStack(n, 1)}
		}
		return []item.Stack{item.NewStack(n, fortuneDiscreteCount(2, 4, 7, enchantments))}
	})
}

// CompostChance ...
func (NetherWart) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (NetherWart) EncodeItem() (name string, meta int16) {
	return "minecraft:nether_wart", 0
}

// EncodeBlock ...
func (n NetherWart) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:nether_wart", map[string]any{"age": int32(n.Age)}
}

// allNetherWart ...
func allNetherWart() (wart []world.Block) {
	for i := 0; i < 4; i++ {
		wart = append(wart, NetherWart{Age: i})
	}
	return
}
