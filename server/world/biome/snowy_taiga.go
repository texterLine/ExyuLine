// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SnowyTaiga ...
type SnowyTaiga struct{}

// Temperature ...
func (SnowyTaiga) Temperature() float64 {
	return -0.5
}

// Rainfall ...
func (SnowyTaiga) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (SnowyTaiga) Depth() float64 {
	return 0.2
}

// Scale ...
func (SnowyTaiga) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (SnowyTaiga) WaterColour() color.RGBA {
	return color.RGBA{R: 0x20, G: 0x5e, B: 0x83, A: 0xa5}
}

// Tags ...
func (SnowyTaiga) Tags() []string {
	return []string{"animal", "cold", "forest", "monster", "overworld", "taiga", "has_structure_trail_ruins", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (SnowyTaiga) String() string {
	return "cold_taiga"
}

// EncodeBiome ...
func (SnowyTaiga) EncodeBiome() int {
	return 30
}
