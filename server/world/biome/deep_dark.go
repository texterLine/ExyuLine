// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DeepDark ...
type DeepDark struct{}

// Temperature ...
func (DeepDark) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (DeepDark) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (DeepDark) Depth() float64 {
	return 0.1
}

// Scale ...
func (DeepDark) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (DeepDark) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (DeepDark) Tags() []string {
	return []string{"caves", "deep_dark", "overworld", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs"}
}

// String ...
func (DeepDark) String() string {
	return "deep_dark"
}

// EncodeBiome ...
func (DeepDark) EncodeBiome() int {
	return 190
}
