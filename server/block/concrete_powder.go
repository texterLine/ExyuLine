// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// ConcretePowder is a gravity affected block that comes in 16 different colours. When interacting with water,
// it becomes concrete.
type ConcretePowder struct {
	gravityAffected
	solid
	snare

	// Colour is the colour of the concrete powder.
	Colour item.Colour
}

// Solidifies ...
func (c ConcretePowder) Solidifies(pos cube.Pos, tx *world.Tx) bool {
	_, water := tx.Block(pos).(Water)
	return water
}

// NeighbourUpdateTick ...
func (c ConcretePowder) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	for i := cube.Face(0); i < 6; i++ {
		if _, ok := tx.Block(pos.Side(i)).(Water); ok {
			tx.SetBlock(pos, Concrete{Colour: c.Colour}, nil)
			return
		}
	}
	c.fall(c, pos, tx)
}

// BreakInfo ...
func (c ConcretePowder) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(c))
}

// EncodeItem ...
func (c ConcretePowder) EncodeItem() (name string, meta int16) {
	return "minecraft:" + c.Colour.String() + "_concrete_powder", 0
}

// EncodeBlock ...
func (c ConcretePowder) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + c.Colour.String() + "_concrete_powder", nil
}

// allConcretePowder returns concrete powder with all possible colours.
func allConcretePowder() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range item.Colours() {
		b = append(b, ConcretePowder{Colour: c})
	}
	return b
}
