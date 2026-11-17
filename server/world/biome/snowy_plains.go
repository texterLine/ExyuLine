// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SnowyPlains ...
type SnowyPlains struct{}

// Temperature ...
func (SnowyPlains) Temperature() float64 {
	return 0
}

// Rainfall ...
func (SnowyPlains) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (SnowyPlains) Depth() float64 {
	return 0.125
}

// Scale ...
func (SnowyPlains) Scale() float64 {
	return 0.05
}

// WaterColour ...
func (SnowyPlains) WaterColour() color.RGBA {
	return color.RGBA{R: 0x14, G: 0x55, B: 0x9b, A: 0xa5}
}

// Tags ...
func (SnowyPlains) Tags() []string {
	return []string{"frozen", "ice", "ice_plains", "overworld", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs", "spawns_snow_foxes", "spawns_white_rabbits"}
}

// String ...
func (SnowyPlains) String() string {
	return "ice_plains"
}

// EncodeBiome ...
func (SnowyPlains) EncodeBiome() int {
	return 12
}
