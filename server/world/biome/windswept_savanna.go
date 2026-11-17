// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// WindsweptSavanna ...
type WindsweptSavanna struct{}

// Temperature ...
func (WindsweptSavanna) Temperature() float64 {
	return 2
}

// Rainfall ...
func (WindsweptSavanna) Rainfall() float64 {
	return 0
}

// Depth ...
func (WindsweptSavanna) Depth() float64 {
	return 0.3625
}

// Scale ...
func (WindsweptSavanna) Scale() float64 {
	return 1.225
}

// WaterColour ...
func (WindsweptSavanna) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (WindsweptSavanna) Tags() []string {
	return []string{"animal", "monster", "mutated", "overworld", "savanna", "spawns_savanna_mobs", "spawns_warm_variant_farm_animals"}
}

// String ...
func (WindsweptSavanna) String() string {
	return "savanna_mutated"
}

// EncodeBiome ...
func (WindsweptSavanna) EncodeBiome() int {
	return 163
}
