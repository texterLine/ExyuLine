// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DripstoneCaves ...
type DripstoneCaves struct{}

// Temperature ...
func (DripstoneCaves) Temperature() float64 {
	return 0.2
}

// Rainfall ...
func (DripstoneCaves) Rainfall() float64 {
	return 0
}

// Depth ...
func (DripstoneCaves) Depth() float64 {
	return 0.1
}

// Scale ...
func (DripstoneCaves) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (DripstoneCaves) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (DripstoneCaves) Tags() []string {
	return []string{"caves", "overworld", "dripstone_caves", "monster"}
}

// String ...
func (DripstoneCaves) String() string {
	return "dripstone_caves"
}

// EncodeBiome ...
func (DripstoneCaves) EncodeBiome() int {
	return 188
}
