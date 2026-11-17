// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// ShatteredSavannaPlateau ...
type ShatteredSavannaPlateau struct{}

// Temperature ...
func (ShatteredSavannaPlateau) Temperature() float64 {
	return 1
}

// Rainfall ...
func (ShatteredSavannaPlateau) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (ShatteredSavannaPlateau) Depth() float64 {
	return 1.05
}

// Scale ...
func (ShatteredSavannaPlateau) Scale() float64 {
	return 1.2125
}

// WaterColour ...
func (ShatteredSavannaPlateau) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (ShatteredSavannaPlateau) Tags() []string {
	return []string{"animal", "monster", "mutated", "overworld", "plateau", "savanna", "spawns_savanna_mobs", "spawns_warm_variant_farm_animals"}
}

// String ...
func (ShatteredSavannaPlateau) String() string {
	return "savanna_plateau_mutated"
}

// EncodeBiome ...
func (ShatteredSavannaPlateau) EncodeBiome() int {
	return 164
}
