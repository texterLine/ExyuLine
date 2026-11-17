// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Plains ...
type Plains struct{}

// Temperature ...
func (Plains) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (Plains) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (Plains) Depth() float64 {
	return 0.125
}

// Scale ...
func (Plains) Scale() float64 {
	return 0.05
}

// WaterColour ...
func (Plains) WaterColour() color.RGBA {
	return color.RGBA{R: 0x44, G: 0xaf, B: 0xf5, A: 0xa5}
}

// Tags ...
func (Plains) Tags() []string {
	return []string{"animal", "monster", "overworld", "plains", "bee_habitat"}
}

// String ...
func (Plains) String() string {
	return "plains"
}

// EncodeBiome ...
func (Plains) EncodeBiome() int {
	return 1
}
