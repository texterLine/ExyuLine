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

// ShelfMushroom is a model used by shelf mushroom blocks. It is a thin plate on the side of the block it is
// attached to, reaching further out once the mushroom has grown.
type ShelfMushroom struct {
	// Facing is the direction the shelf mushroom faces, away from the block it is attached to.
	Facing cube.Direction
	// Large is true if the shelf mushroom has grown to its large stage.
	Large bool
}

// BBox returns a plate hugging the side behind the mushroom.
func (s ShelfMushroom) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	depth := 0.5
	if s.Large {
		depth = 0.75
	}
	return []cube.BBox{full.
		ExtendTowards(cube.FaceUp, -0.75).
		ExtendTowards(s.Facing.Face(), -(1 - depth))}
}

// FaceSolid always returns false.
func (ShelfMushroom) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
