// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// Barrier is a transparent solid block used to create invisible boundaries.
type Barrier struct {
	sourceWaterDisplacer
	transparent
	solid
}

// SideClosed ...
func (Barrier) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// EncodeItem ...
func (Barrier) EncodeItem() (name string, meta int16) {
	return "minecraft:barrier", 0
}

// EncodeBlock ...
func (Barrier) EncodeBlock() (string, map[string]any) {
	return "minecraft:barrier", nil
}
