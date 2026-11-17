// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// OldGrowthBirchForest ...
type OldGrowthBirchForest struct{}

// Temperature ...
func (OldGrowthBirchForest) Temperature() float64 {
	return 0.6
}

// Rainfall ...
func (OldGrowthBirchForest) Rainfall() float64 {
	return 0.6
}

// Depth ...
func (OldGrowthBirchForest) Depth() float64 {
	return 0.2
}

// Scale ...
func (OldGrowthBirchForest) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (OldGrowthBirchForest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x06, G: 0x77, B: 0xce, A: 0xa5}
}

// Tags ...
func (OldGrowthBirchForest) Tags() []string {
	return []string{"animal", "birch", "forest", "monster", "mutated", "bee_habitat", "overworld_generation", "has_structure_trail_ruins", "has_structure_abandoned_camp"}
}

// String ...
func (OldGrowthBirchForest) String() string {
	return "birch_forest_mutated"
}

// EncodeBiome ...
func (OldGrowthBirchForest) EncodeBiome() int {
	return 155
}
