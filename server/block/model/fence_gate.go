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

// FenceGate is a model used by fence gates. The model is completely zero-ed when the FenceGate is opened.
type FenceGate struct {
	// Facing is the facing direction of the FenceGate. A fence gate can only be opened in this direction or the
	// direction opposite to it.
	Facing cube.Direction
	// Open specifies if the FenceGate is open. In this case, BBox returns an empty slice.
	Open bool
}

// BBox returns up to one physics.BBox depending on the facing direction of the FenceGate and whether it is open.
func (f FenceGate) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	if f.Open {
		return nil
	}
	return []cube.BBox{cube.Box(0, 0, 0, 1, 1.5, 1).Stretch(f.Facing.Face().Axis(), -0.375)}
}

// FaceSolid always returns false.
func (f FenceGate) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
