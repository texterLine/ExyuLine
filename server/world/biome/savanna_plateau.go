// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SavannaPlateau ...
type SavannaPlateau struct{}

// Temperature ...
func (SavannaPlateau) Temperature() float64 {
	return 1
}

// Rainfall ...
func (SavannaPlateau) Rainfall() float64 {
	return 0
}

// Depth ...
func (SavannaPlateau) Depth() float64 {
	return 1.5
}

// Scale ...
func (SavannaPlateau) Scale() float64 {
	return 0.025
}

// WaterColour ...
func (SavannaPlateau) WaterColour() color.RGBA {
	return color.RGBA{R: 0x25, G: 0x90, B: 0xa8, A: 0xa5}
}

// Tags ...
func (SavannaPlateau) Tags() []string {
	return []string{"animal", "monster", "overworld", "plateau", "savanna", "spawns_savanna_mobs", "spawns_warm_variant_farm_animals"}
}

// String ...
func (SavannaPlateau) String() string {
	return "savanna_plateau"
}

// EncodeBiome ...
func (SavannaPlateau) EncodeBiome() int {
	return 36
}
