// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SnowyBeach ...
type SnowyBeach struct{}

// Temperature ...
func (SnowyBeach) Temperature() float64 {
	return 0.05
}

// Rainfall ...
func (SnowyBeach) Rainfall() float64 {
	return 0.3
}

// Depth ...
func (SnowyBeach) Depth() float64 {
	return 0
}

// Scale ...
func (SnowyBeach) Scale() float64 {
	return 0.025
}

// WaterColour ...
func (SnowyBeach) WaterColour() color.RGBA {
	return color.RGBA{R: 0x14, G: 0x63, B: 0xa5, A: 0xa5}
}

// Tags ...
func (SnowyBeach) Tags() []string {
	return []string{"beach", "cold", "monster", "overworld", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs", "spawns_snow_foxes", "spawns_white_rabbits"}
}

// String ...
func (SnowyBeach) String() string {
	return "cold_beach"
}

// EncodeBiome ...
func (SnowyBeach) EncodeBiome() int {
	return 26
}
