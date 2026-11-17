// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// ModifiedWoodedBadlandsPlateau ...
type ModifiedWoodedBadlandsPlateau struct{}

// Temperature ...
func (ModifiedWoodedBadlandsPlateau) Temperature() float64 {
	return 2
}

// Rainfall ...
func (ModifiedWoodedBadlandsPlateau) Rainfall() float64 {
	return 0
}

// Depth ...
func (ModifiedWoodedBadlandsPlateau) Depth() float64 {
	return 0.45
}

// Scale ...
func (ModifiedWoodedBadlandsPlateau) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (ModifiedWoodedBadlandsPlateau) WaterColour() color.RGBA {
	return color.RGBA{R: 0x55, G: 0x80, B: 0x9e, A: 0xa5}
}

// Tags ...
func (ModifiedWoodedBadlandsPlateau) Tags() []string {
	return []string{"animal", "mesa", "monster", "mutated", "overworld", "plateau", "spawns_mesa_mobs", "spawns_warm_variant_farm_animals", "surface_mineshaft"}
}

// String ...
func (ModifiedWoodedBadlandsPlateau) String() string {
	return "mesa_plateau_stone_mutated"
}

// EncodeBiome ...
func (ModifiedWoodedBadlandsPlateau) EncodeBiome() int {
	return 166
}
