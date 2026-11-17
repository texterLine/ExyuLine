// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

import (
	"github.com/texterLine/exyuline/server/block/cube"
)

// BlockModel represents the model of a block. These models specify the ways a block can be collided with and
// whether specific faces are solid wrt. being able to, for example, place torches onto those sides.
type BlockModel interface {
	// BBox returns the bounding boxes that a block with this model can be collided with.
	BBox(pos cube.Pos, s BlockSource) []cube.BBox
	// FaceSolid checks if a specific face of a block at the position in a world passed is solid. Blocks may
	// be attached to these faces.
	FaceSolid(pos cube.Pos, face cube.Face, s BlockSource) bool
}

// unknownModel is the model used for unknown blocks. It is the equivalent of a fully solid model.
type unknownModel struct{}

// BBox ...
func (u unknownModel) BBox(cube.Pos, BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0, 0, 0, 1, 1, 1)}
}

// FaceSolid ...
func (u unknownModel) FaceSolid(cube.Pos, cube.Face, BlockSource) bool {
	return true
}
