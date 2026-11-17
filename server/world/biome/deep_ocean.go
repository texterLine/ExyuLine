// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DeepOcean ...
type DeepOcean struct{}

// Temperature ...
func (DeepOcean) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (DeepOcean) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (DeepOcean) Depth() float64 {
	return -1.8
}

// Scale ...
func (DeepOcean) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (DeepOcean) WaterColour() color.RGBA {
	return color.RGBA{R: 0x17, G: 0x87, B: 0xd4, A: 0xa5}
}

// Tags ...
func (DeepOcean) Tags() []string {
	return []string{"deep", "monster", "ocean", "overworld", "spawns_warm_variant_farm_animals", "fast_fishing", "high_seas"}
}

// String ...
func (DeepOcean) String() string {
	return "deep_ocean"
}

// EncodeBiome ...
func (DeepOcean) EncodeBiome() int {
	return 24
}
