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

// Campfire is the model used by campfires.
type Campfire struct{}

// BBox returns a flat BBox with a height of 0.4375.
func (Campfire) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0, 0, 0, 1, 0.4375, 1)}
}

// FaceSolid returns true if the face is down.
func (Campfire) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	return face == cube.FaceDown
}
