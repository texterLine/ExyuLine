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

// Chain is a model used by chain blocks.
type Chain struct {
	// Axis is the axis which the chain faces.
	Axis cube.Axis
}

// BBox ...
func (c Chain) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{cube.Box(0.40625, 0.40625, 0.40625, 0.59375, 0.59375, 0.59375).Stretch(c.Axis, 0.40625)}
}

// FaceSolid ...
func (Chain) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
