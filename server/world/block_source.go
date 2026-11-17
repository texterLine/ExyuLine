// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

import "github.com/texterLine/exyuline/server/block/cube"

// BlockSource represents a source for obtaining blocks.
type BlockSource interface {
	// Block returns the block at the given position in the block source.
	Block(cube.Pos) Block
}

// StateDeriver is implemented by Blocks that derive part of their state from the blocks around them,
// such as the connections of a fence. World uses it to recalculate that state for chunks saved before
// it existed, which hold a default in its place.
type StateDeriver interface {
	Block
	// DeriveState returns the Block with its derived state recalculated against src.
	DeriveState(pos cube.Pos, src BlockSource) Block
}

// worldSource is a wrapper around a world transaction that implements BlockSource.
type worldSource struct{ tx *Tx }

func (w worldSource) Block(pos cube.Pos) Block { return w.tx.block(pos) }
