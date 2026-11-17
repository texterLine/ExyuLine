// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// BasaltDeltas ...
type BasaltDeltas struct{}

// Temperature ...
func (BasaltDeltas) Temperature() float64 {
	return 2
}

// Rainfall ...
func (BasaltDeltas) Rainfall() float64 {
	return 0
}

// Depth ...
func (BasaltDeltas) Depth() float64 {
	return 0.1
}

// Scale ...
func (BasaltDeltas) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (BasaltDeltas) WaterColour() color.RGBA {
	return color.RGBA{R: 0x3f, G: 0x76, B: 0xe4, A: 0xa5}
}

// Tags ...
func (BasaltDeltas) Tags() []string {
	return []string{"nether", "basalt_deltas", "spawn_many_magma_cubes", "spawn_ghast", "spawns_warm_variant_farm_animals"}
}

// String ...
func (BasaltDeltas) String() string {
	return "basalt_deltas"
}

// EncodeBiome ...
func (BasaltDeltas) EncodeBiome() int {
	return 181
}
