// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// WindsweptGravellyHills ...
type WindsweptGravellyHills struct{}

// Temperature ...
func (WindsweptGravellyHills) Temperature() float64 {
	return 0.2
}

// Rainfall ...
func (WindsweptGravellyHills) Rainfall() float64 {
	return 0.3
}

// Depth ...
func (WindsweptGravellyHills) Depth() float64 {
	return 1
}

// Scale ...
func (WindsweptGravellyHills) Scale() float64 {
	return 0.5
}

// WaterColour ...
func (WindsweptGravellyHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0e, G: 0x63, B: 0xab, A: 0xa5}
}

// Tags ...
func (WindsweptGravellyHills) Tags() []string {
	return []string{"animal", "extreme_hills", "monster", "mutated", "overworld", "spawns_cold_variant_farm_animals"}
}

// String ...
func (WindsweptGravellyHills) String() string {
	return "extreme_hills_mutated"
}

// EncodeBiome ...
func (WindsweptGravellyHills) EncodeBiome() int {
	return 131
}
