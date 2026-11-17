// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// River ...
type River struct{}

// Temperature ...
func (River) Temperature() float64 {
	return 0.5
}

// Rainfall ...
func (River) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (River) Depth() float64 {
	return -0.5
}

// Scale ...
func (River) Scale() float64 {
	return 0
}

// WaterColour ...
func (River) WaterColour() color.RGBA {
	return color.RGBA{R: 0x00, G: 0x84, B: 0xff, A: 0xa5}
}

// Tags ...
func (River) Tags() []string {
	return []string{"overworld", "spawns_more_frequent_drowned", "spawns_reduced_water_ambient_mobs", "spawns_river_mobs", "river"}
}

// String ...
func (River) String() string {
	return "river"
}

// EncodeBiome ...
func (River) EncodeBiome() int {
	return 7
}
