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

// Portal is a model used by portal blocks.
type Portal struct {
	// Axis is the axis that the portal faces.
	Axis cube.Axis
}

// BBox ...
func (Portal) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return nil
}

// FaceSolid ...
func (Portal) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
