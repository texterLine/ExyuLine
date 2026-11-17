// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/world/sound"
)

// PackedIce is an opaque solid block variant of ice. Unlike regular ice, it does not melt near bright light sources.
type PackedIce struct {
	solid
}

// Instrument ...
func (PackedIce) Instrument() sound.Instrument {
	return sound.Chimes()
}

// BreakInfo ...
func (p PackedIce) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, pickaxeEffective, silkTouchOnlyDrop(p))
}

// Friction ...
func (p PackedIce) Friction() float64 {
	return 0.98
}

// EncodeItem ...
func (PackedIce) EncodeItem() (name string, meta int16) {
	return "minecraft:packed_ice", 0
}

// EncodeBlock ...
func (PackedIce) EncodeBlock() (string, map[string]any) {
	return "minecraft:packed_ice", nil
}
