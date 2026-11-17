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

// TilledGrass is a model used for grass that has been tilled in some way, such as dirt paths and farmland.
type TilledGrass struct{}

// BBox returns a physics.BBox that spans an entire block.
func (TilledGrass) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{full.ExtendTowards(cube.FaceUp, -0.0625)}
}

// FaceSolid always returns true.
func (TilledGrass) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return true
}
