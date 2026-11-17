// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// BadlandsPlateau ...
type BadlandsPlateau struct{}

// Temperature ...
func (BadlandsPlateau) Temperature() float64 {
	return 2
}

// Rainfall ...
func (BadlandsPlateau) Rainfall() float64 {
	return 0
}

// Depth ...
func (BadlandsPlateau) Depth() float64 {
	return 1.5
}

// Scale ...
func (BadlandsPlateau) Scale() float64 {
	return 0.025
}

// WaterColour ...
func (BadlandsPlateau) WaterColour() color.RGBA {
	return color.RGBA{R: 0x55, G: 0x80, B: 0x9e, A: 0xa5}
}

// Tags ...
func (BadlandsPlateau) Tags() []string {
	return []string{"animal", "mesa", "monster", "overworld", "plateau", "rare", "spawns_mesa_mobs", "spawns_warm_variant_farm_animals", "surface_mineshaft"}
}

// String ...
func (BadlandsPlateau) String() string {
	return "mesa_plateau"
}

// EncodeBiome ...
func (BadlandsPlateau) EncodeBiome() int {
	return 39
}
