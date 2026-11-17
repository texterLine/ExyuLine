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

// Bamboo is a model used by bamboo.
type Bamboo struct {
	Thick bool
}

// BBox ...
func (b Bamboo) BBox(pos cube.Pos, s world.BlockSource) []cube.BBox {
	// The stalk's box extends from the block's centre towards positive X and Z.
	size := 0.5 + 2.0/16.0
	if b.Thick {
		size = 0.5 + 3.0/16.0
	}
	offset := randomOffset(pos, -0.25, 0.25, 16)
	return []cube.BBox{cube.Box(0.5, 0, 0.5, size, 1, size).Translate(offset)}
}

// FaceSolid ...
func (b Bamboo) FaceSolid(pos cube.Pos, face cube.Face, s world.BlockSource) bool {
	return false
}
