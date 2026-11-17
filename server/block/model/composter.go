// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package model

import (
	"math"

	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// Composter is a model used by composter blocks. It is solid on all sides apart from the top, and the height of the
// inside area depends on the level of compost inside the composter.
type Composter struct {
	// Level is the level of compost inside the composter.
	Level int
}

// BBox ...
func (c Composter) BBox(_ cube.Pos, _ world.BlockSource) []cube.BBox {
	compostHeight := math.Abs(math.Min(float64(c.Level), 7)*0.125 - 0.0625)
	return []cube.BBox{
		cube.Box(0, 0, 0, 1, 1, 0.125),
		cube.Box(0, 0, 0.875, 1, 1, 1),
		cube.Box(0.875, 0, 0, 1, 1, 1),
		cube.Box(0, 0, 0, 0.125, 1, 1),
		cube.Box(0.125, 0, 0.125, 0.875, 0.125+compostHeight, 0.875),
	}
}

// FaceSolid returns true for all faces other than the top.
func (Composter) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	return face != cube.FaceUp
}
