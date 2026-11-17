// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// OldGrowthSpruceTaiga ...
type OldGrowthSpruceTaiga struct{}

// Temperature ...
func (OldGrowthSpruceTaiga) Temperature() float64 {
	return 0.25
}

// Rainfall ...
func (OldGrowthSpruceTaiga) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (OldGrowthSpruceTaiga) Depth() float64 {
	return 0.2
}

// Scale ...
func (OldGrowthSpruceTaiga) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (OldGrowthSpruceTaiga) WaterColour() color.RGBA {
	return color.RGBA{R: 0x2d, G: 0x6d, B: 0x77, A: 0xa5}
}

// Tags ...
func (OldGrowthSpruceTaiga) Tags() []string {
	return []string{"animal", "forest", "mega", "monster", "mutated", "overworld", "taiga", "has_structure_trail_ruins", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (OldGrowthSpruceTaiga) String() string {
	return "redwood_taiga_mutated"
}

// EncodeBiome ...
func (OldGrowthSpruceTaiga) EncodeBiome() int {
	return 160
}
