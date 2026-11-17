// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SunflowerPlains ...
type SunflowerPlains struct{}

// Temperature ...
func (SunflowerPlains) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (SunflowerPlains) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (SunflowerPlains) Depth() float64 {
	return 0.125
}

// Scale ...
func (SunflowerPlains) Scale() float64 {
	return 0.05
}

// WaterColour ...
func (SunflowerPlains) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (SunflowerPlains) Tags() []string {
	return []string{"animal", "monster", "mutated", "overworld", "plains", "bee_habitat"}
}

// String ...
func (SunflowerPlains) String() string {
	return "sunflower_plains"
}

// EncodeBiome ...
func (SunflowerPlains) EncodeBiome() int {
	return 129
}
