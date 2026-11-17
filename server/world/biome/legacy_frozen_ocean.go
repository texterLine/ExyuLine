// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// LegacyFrozenOcean ...
type LegacyFrozenOcean struct{}

// Temperature ...
func (LegacyFrozenOcean) Temperature() float64 {
	return 0
}

// Rainfall ...
func (LegacyFrozenOcean) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (LegacyFrozenOcean) Depth() float64 {
	return -1
}

// Scale ...
func (LegacyFrozenOcean) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (LegacyFrozenOcean) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (LegacyFrozenOcean) Tags() []string {
	return []string{"legacy", "frozen", "ocean", "overworld", "spawns_cold_variant_farm_animals", "spawns_polar_bears_on_alternate_blocks", "fast_fishing", "high_seas"}
}

// String ...
func (LegacyFrozenOcean) String() string {
	return "legacy_frozen_ocean"
}

// EncodeBiome ...
func (LegacyFrozenOcean) EncodeBiome() int {
	return 10
}
