// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SulfurCaves ...
type SulfurCaves struct{}

// Temperature ...
func (SulfurCaves) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (SulfurCaves) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (SulfurCaves) Depth() float64 {
	return 0.1
}

// Scale ...
func (SulfurCaves) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (SulfurCaves) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (SulfurCaves) Tags() []string {
	return []string{"caves", "sulfur_caves", "overworld", "monster"}
}

// String ...
func (SulfurCaves) String() string {
	return "sulfur_caves"
}

// EncodeBiome ...
func (SulfurCaves) EncodeBiome() int {
	return 194
}
