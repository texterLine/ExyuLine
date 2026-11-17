// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// TallBirchHills ...
type TallBirchHills struct{}

// Temperature ...
func (TallBirchHills) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (TallBirchHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (TallBirchHills) Depth() float64 {
	return 0.55
}

// Scale ...
func (TallBirchHills) Scale() float64 {
	return 0.5
}

// WaterColour ...
func (TallBirchHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0a, G: 0x74, B: 0xc4, A: 0xa5}
}

// Tags ...
func (TallBirchHills) Tags() []string {
	return []string{"animal", "birch", "forest", "hills", "monster", "mutated", "overworld_generation"}
}

// String ...
func (TallBirchHills) String() string {
	return "birch_forest_hills_mutated"
}

// EncodeBiome ...
func (TallBirchHills) EncodeBiome() int {
	return 156
}
