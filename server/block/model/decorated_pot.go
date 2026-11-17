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

// DecoratedPot is the model for a DecoratedPot. It is just barely not a full block, having a slightly reduced width and
// depth.
type DecoratedPot struct{}

// BBox returns a physics.BBox that is slightly smaller than a full block.
func (DecoratedPot) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0.025, 0, 0.025, 0.975, 1, 0.975)}
}

// FaceSolid always returns true for the top and bottom face, and false for all other faces.
func (DecoratedPot) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	return face.Axis() == cube.Y
}
