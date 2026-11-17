// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// WindsweptForest ...
type WindsweptForest struct{}

// Temperature ...
func (WindsweptForest) Temperature() float64 {
	return 0.2
}

// Rainfall ...
func (WindsweptForest) Rainfall() float64 {
	return 0.3
}

// Depth ...
func (WindsweptForest) Depth() float64 {
	return 1
}

// Scale ...
func (WindsweptForest) Scale() float64 {
	return 0.5
}

// WaterColour ...
func (WindsweptForest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0e, G: 0x63, B: 0xab, A: 0xa5}
}

// Tags ...
func (WindsweptForest) Tags() []string {
	return []string{"animal", "extreme_hills", "forest", "monster", "mountain", "overworld", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (WindsweptForest) String() string {
	return "extreme_hills_plus_trees"
}

// EncodeBiome ...
func (WindsweptForest) EncodeBiome() int {
	return 34
}
