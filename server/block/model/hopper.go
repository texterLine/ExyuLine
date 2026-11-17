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

// Hopper is a model used by hoppers.
type Hopper struct{}

// BBox returns a physics.BBox that spans a full block.
func (h Hopper) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	bbox := []cube.BBox{full.ExtendTowards(cube.FaceUp, -0.375)}
	for _, f := range cube.HorizontalFaces() {
		bbox = append(bbox, full.ExtendTowards(f, -0.875))
	}
	return bbox
}

// FaceSolid only returns true for the top face of the hopper.
func (Hopper) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	return face == cube.FaceUp
}
