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

// Skull is a model used by skull blocks.
type Skull struct {
	// Direction is the direction the skull is facing.
	Direction cube.Face
	// Hanging specifies if the Skull is hanging on a wall.
	Hanging bool
}

// BBox ...
func (s Skull) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	box := cube.Box(0.25, 0, 0.25, 0.75, 0.5, 0.75)
	if !s.Hanging {
		return []cube.BBox{box}
	}
	return []cube.BBox{box.TranslateTowards(s.Direction.Opposite(), 0.25).TranslateTowards(cube.FaceUp, 0.25)}
}

// FaceSolid ...
func (Skull) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
