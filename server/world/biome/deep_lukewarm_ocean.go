// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DeepLukewarmOcean ...
type DeepLukewarmOcean struct{}

// Temperature ...
func (DeepLukewarmOcean) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (DeepLukewarmOcean) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (DeepLukewarmOcean) Depth() float64 {
	return -1.8
}

// Scale ...
func (DeepLukewarmOcean) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (DeepLukewarmOcean) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0d, G: 0x96, B: 0xdb, A: 0xa5}
}

// Tags ...
func (DeepLukewarmOcean) Tags() []string {
	return []string{"deep", "lukewarm", "monster", "ocean", "overworld", "spawns_warm_variant_farm_animals", "fast_fishing", "high_seas", "temperate_ocean"}
}

// String ...
func (DeepLukewarmOcean) String() string {
	return "deep_lukewarm_ocean"
}

// EncodeBiome ...
func (DeepLukewarmOcean) EncodeBiome() int {
	return 43
}
