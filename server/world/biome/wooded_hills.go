// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// WoodedHills ...
type WoodedHills struct{}

// Temperature ...
func (WoodedHills) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (WoodedHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (WoodedHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (WoodedHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (WoodedHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x05, G: 0x6b, B: 0xd1, A: 0xa5}
}

// Tags ...
func (WoodedHills) Tags() []string {
	return []string{"animal", "hills", "monster", "overworld", "forest", "bee_habitat"}
}

// String ...
func (WoodedHills) String() string {
	return "forest_hills"
}

// EncodeBiome ...
func (WoodedHills) EncodeBiome() int {
	return 18
}
