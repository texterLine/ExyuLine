// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Meadow ...
type Meadow struct{}

// Temperature ...
func (Meadow) Temperature() float64 {
	return 0.3
}

// Rainfall ...
func (Meadow) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (Meadow) Depth() float64 {
	return 0.1
}

// Scale ...
func (Meadow) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (Meadow) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (Meadow) Tags() []string {
	return []string{"mountains", "monster", "overworld", "meadow", "bee_habitat", "has_structure_abandoned_camp"}
}

// String ...
func (Meadow) String() string {
	return "meadow"
}

// EncodeBiome ...
func (Meadow) EncodeBiome() int {
	return 186
}
