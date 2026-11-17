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

// Cake is a model used by cake blocks.
type Cake struct {
	// Bites is the amount of bites that were taken from the cake. A cake can have up to 7 bites taken from it, before
	// being consumed entirely.
	Bites int
}

// BBox returns an BBox with a size that depends on the amount of bites taken.
func (c Cake) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0.0625, 0, 0.0625, 0.9375, 0.5, 0.9375).
		ExtendTowards(cube.FaceWest, -(float64(c.Bites) / 8))}
}

// FaceSolid always returns false.
func (c Cake) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
