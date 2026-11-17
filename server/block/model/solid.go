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

// Solid is the model of a fully solid block. Blocks with this model, such as stone or wooden planks, have a
// 1x1x1 collision box.
type Solid struct{}

// full is a BBox that occupies a full block.
var full = cube.Box(0, 0, 0, 1, 1, 1)

// BBox returns a physics.BBox spanning a full block.
func (Solid) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{full}
}

// FaceSolid always returns true.
func (Solid) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return true
}
