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

// Lantern is a model for the lantern block. It can be placed on the ground or hanging from the ceiling.
type Lantern struct {
	// Hanging specifies if the lantern is hanging from a block or if it's placed on the ground.
	Hanging bool
}

// BBox returns a physics.BBox attached to either the ceiling or to the ground.
func (l Lantern) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	if l.Hanging {
		return []cube.BBox{cube.Box(0.3125, 0.125, 0.3125, 0.6875, 0.625, 0.6875)}
	}
	return []cube.BBox{cube.Box(0.3125, 0, 0.3125, 0.6875, 0.5, 0.6875)}
}

// FaceSolid always returns false.
func (l Lantern) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
