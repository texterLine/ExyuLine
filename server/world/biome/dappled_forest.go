// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DappledForest ...
type DappledForest struct{}

// Temperature ...
func (DappledForest) Temperature() float64 {
	return 0.6
}

// Rainfall ...
func (DappledForest) Rainfall() float64 {
	return 0.6
}

// Depth ...
func (DappledForest) Depth() float64 {
	return 0.1
}

// Scale ...
func (DappledForest) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (DappledForest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (DappledForest) Tags() []string {
	return []string{"animal", "dappled_forest", "monster", "forest", "overworld", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (DappledForest) String() string {
	return "dappled_forest"
}

// EncodeBiome ...
func (DappledForest) EncodeBiome() int {
	return 195
}
