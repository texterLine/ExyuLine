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

// EndRod is a model used by end rod blocks.
type EndRod struct {
	// Axis is the axis which the end rod faces.
	Axis cube.Axis
}

// BBox ...
func (e EndRod) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0.375, 0.375, 0.375, 0.625, 0.625, 0.625).Stretch(e.Axis, 0.375)}
}

// FaceSolid ...
func (EndRod) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
