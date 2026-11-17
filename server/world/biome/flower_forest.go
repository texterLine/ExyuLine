// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// FlowerForest ...
type FlowerForest struct{}

// Temperature ...
func (FlowerForest) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (FlowerForest) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (FlowerForest) Depth() float64 {
	return 0.1
}

// Scale ...
func (FlowerForest) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (FlowerForest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x20, G: 0xa3, B: 0xcc, A: 0xa5}
}

// Tags ...
func (FlowerForest) Tags() []string {
	return []string{"animal", "flower_forest", "monster", "mutated", "overworld", "bee_habitat", "has_structure_abandoned_camp"}
}

// String ...
func (FlowerForest) String() string {
	return "flower_forest"
}

// EncodeBiome ...
func (FlowerForest) EncodeBiome() int {
	return 132
}
