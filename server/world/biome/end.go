// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// End ...
type End struct{}

// Temperature ...
func (End) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (End) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (End) Depth() float64 {
	return 0.1
}

// Scale ...
func (End) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (End) WaterColour() color.RGBA {
	return color.RGBA{R: 0x62, G: 0x52, B: 0x9e, A: 0xa5}
}

// Tags ...
func (End) Tags() []string {
	return []string{"the_end", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs"}
}

// String ...
func (End) String() string {
	return "the_end"
}

// EncodeBiome ...
func (End) EncodeBiome() int {
	return 9
}
