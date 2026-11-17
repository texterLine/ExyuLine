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

// Trapdoor is a model used for trapdoors. It has no solid faces and a bounding box that changes depending on
// the direction of the trapdoor.
type Trapdoor struct {
	// Facing is the facing direction of the Trapdoor. In addition to the texture, it influences the direction in which
	// the Trapdoor is opened.
	Facing cube.Direction
	// Open and Top specify if the Trapdoor is opened and if it's in the top or bottom part of a block respectively.
	Open, Top bool
}

// BBox returns a physics.BBox that depends on the facing direction of the Trapdoor and whether it is open and in the
// top part of the block.
func (t Trapdoor) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	if t.Open {
		return []cube.BBox{full.ExtendTowards(t.Facing.Face(), -0.8125)}
	} else if t.Top {
		return []cube.BBox{cube.Box(0, 0.8125, 0, 1, 1, 1)}
	}
	return []cube.BBox{cube.Box(0, 0, 0, 1, 0.1875, 1)}
}

// FaceSolid returns true if the face is completely filled with the trapdoor.
func (t Trapdoor) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	if t.Open {
		return t.Facing.Face().Opposite() == face
	} else if t.Top {
		return face == cube.FaceUp
	}
	return face == cube.FaceDown
}
