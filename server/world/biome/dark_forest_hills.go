// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DarkForestHills ...
type DarkForestHills struct{}

// Temperature ...
func (DarkForestHills) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (DarkForestHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (DarkForestHills) Depth() float64 {
	return 0.2
}

// Scale ...
func (DarkForestHills) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (DarkForestHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x3b, G: 0x6c, B: 0xd1, A: 0xa5}
}

// Tags ...
func (DarkForestHills) Tags() []string {
	return []string{"animal", "forest", "monster", "mutated", "roofed", "overworld_generation"}
}

// String ...
func (DarkForestHills) String() string {
	return "roofed_forest_mutated"
}

// EncodeBiome ...
func (DarkForestHills) EncodeBiome() int {
	return 157
}
