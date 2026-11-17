// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// CherryGrove ...
type CherryGrove struct{}

// Temperature ...
func (CherryGrove) Temperature() float64 {
	return 0.3
}

// Rainfall ...
func (CherryGrove) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (CherryGrove) Depth() float64 {
	return 0.1
}

// Scale ...
func (CherryGrove) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (CherryGrove) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (CherryGrove) Tags() []string {
	return []string{"mountains", "monster", "overworld", "cherry_grove", "bee_habitat", "has_structure_abandoned_camp"}
}

// String ...
func (CherryGrove) String() string {
	return "cherry_grove"
}

// EncodeBiome ...
func (CherryGrove) EncodeBiome() int {
	return 192
}
