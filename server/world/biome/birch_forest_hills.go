// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// BirchForestHills ...
type BirchForestHills struct{}

// Temperature ...
func (BirchForestHills) Temperature() float64 {
	return 0.6
}

// Rainfall ...
func (BirchForestHills) Rainfall() float64 {
	return 0.6
}

// Depth ...
func (BirchForestHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (BirchForestHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (BirchForestHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0a, G: 0x74, B: 0xc4, A: 0xa5}
}

// Tags ...
func (BirchForestHills) Tags() []string {
	return []string{"animal", "birch", "forest", "hills", "monster", "overworld", "bee_habitat"}
}

// String ...
func (BirchForestHills) String() string {
	return "birch_forest_hills"
}

// EncodeBiome ...
func (BirchForestHills) EncodeBiome() int {
	return 28
}
