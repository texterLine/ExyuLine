// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// ColdOcean ...
type ColdOcean struct{}

// Temperature ...
func (ColdOcean) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (ColdOcean) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (ColdOcean) Depth() float64 {
	return -1
}

// Scale ...
func (ColdOcean) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (ColdOcean) WaterColour() color.RGBA {
	return color.RGBA{R: 0x20, G: 0x80, B: 0xc9, A: 0xa5}
}

// Tags ...
func (ColdOcean) Tags() []string {
	return []string{"cold", "monster", "ocean", "overworld", "spawns_cold_variant_farm_animals", "fast_fishing", "high_seas"}
}

// String ...
func (ColdOcean) String() string {
	return "cold_ocean"
}

// EncodeBiome ...
func (ColdOcean) EncodeBiome() int {
	return 44
}
