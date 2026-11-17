// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Beach ...
type Beach struct{}

// Temperature ...
func (Beach) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (Beach) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (Beach) Depth() float64 {
	return 0
}

// Scale ...
func (Beach) Scale() float64 {
	return 0.025
}

// WaterColour ...
func (Beach) WaterColour() color.RGBA {
	return color.RGBA{R: 0x15, G: 0x7c, B: 0xab, A: 0xa5}
}

// Tags ...
func (Beach) Tags() []string {
	return []string{"beach", "monster", "overworld", "warm"}
}

// String ...
func (Beach) String() string {
	return "beach"
}

// EncodeBiome ...
func (Beach) EncodeBiome() int {
	return 16
}
