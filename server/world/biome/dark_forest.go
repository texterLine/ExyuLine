// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DarkForest ...
type DarkForest struct{}

// Temperature ...
func (DarkForest) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (DarkForest) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (DarkForest) Depth() float64 {
	return 0.1
}

// Scale ...
func (DarkForest) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (DarkForest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x3b, G: 0x6c, B: 0xd1, A: 0xa5}
}

// Tags ...
func (DarkForest) Tags() []string {
	return []string{"animal", "forest", "monster", "no_legacy_worldgen", "overworld", "roofed"}
}

// String ...
func (DarkForest) String() string {
	return "roofed_forest"
}

// EncodeBiome ...
func (DarkForest) EncodeBiome() int {
	return 29
}
