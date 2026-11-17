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

// Shulker is the model of a shulker box. The bounding box grows along the
// facing axis as the lid opens.
type Shulker struct {
	// Facing is the direction that the lid opens towards.
	Facing cube.Face
	// Progress is the lid animation progress, ranging from 0 (closed) to 10 (fully open).
	Progress int32
}

// BBox returns a single bounding box that extends outward along Facing as the
// lid opens.
func (s Shulker) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	peak := ShulkerPhysicalPeak(s.Progress)
	return []cube.BBox{full.ExtendTowards(s.Facing, peak)}
}

// ShulkerPhysicalPeak returns the lid extension along the facing axis for a
// given Progress in [0, 10]. The curve eases out cubically so the lid moves
// quickly and settles.
func ShulkerPhysicalPeak(progress int32) float64 {
	t := float64(progress) / 10.0
	return (1.0 - (1.0-t)*(1.0-t)*(1.0-t)) * 0.5
}

// FaceSolid always returns false.
func (Shulker) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
