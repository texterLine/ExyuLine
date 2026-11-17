// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// OldGrowthPineTaiga ...
type OldGrowthPineTaiga struct{}

// Temperature ...
func (OldGrowthPineTaiga) Temperature() float64 {
	return 0.3
}

// Rainfall ...
func (OldGrowthPineTaiga) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (OldGrowthPineTaiga) Depth() float64 {
	return 0.2
}

// Scale ...
func (OldGrowthPineTaiga) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (OldGrowthPineTaiga) WaterColour() color.RGBA {
	return color.RGBA{R: 0x2d, G: 0x6d, B: 0x77, A: 0xa5}
}

// Tags ...
func (OldGrowthPineTaiga) Tags() []string {
	return []string{"animal", "forest", "mega", "monster", "overworld", "rare", "taiga", "has_structure_trail_ruins", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (OldGrowthPineTaiga) String() string {
	return "mega_taiga"
}

// EncodeBiome ...
func (OldGrowthPineTaiga) EncodeBiome() int {
	return 32
}
