// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DeepWarmOcean ...
type DeepWarmOcean struct{}

// Temperature ...
func (DeepWarmOcean) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (DeepWarmOcean) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (DeepWarmOcean) Depth() float64 {
	return -1.8
}

// Scale ...
func (DeepWarmOcean) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (DeepWarmOcean) WaterColour() color.RGBA {
	return color.RGBA{R: 0x02, G: 0xb0, B: 0xe5, A: 0xa5}
}

// Tags ...
func (DeepWarmOcean) Tags() []string {
	return []string{"deep", "monster", "ocean", "overworld", "warm", "spawns_warm_variant_farm_animals", "fast_fishing", "high_seas", "temperate_ocean"}
}

// String ...
func (DeepWarmOcean) String() string {
	return "deep_warm_ocean"
}

// EncodeBiome ...
func (DeepWarmOcean) EncodeBiome() int {
	return 41
}
