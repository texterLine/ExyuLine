// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package model

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// Ladder is the model for a ladder block.
type Ladder struct {
	// Facing is the side opposite to the block the Ladder is currently attached to.
	Facing cube.Direction
}

// BBox returns one physics.BBox that depends on the facing direction of the Ladder.
func (l Ladder) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{full.ExtendTowards(l.Facing.Face(), -0.8125)}
}

// FaceSolid always returns false.
func (l Ladder) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
