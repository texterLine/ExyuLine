// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SnowyMountains ...
type SnowyMountains struct{}

// Temperature ...
func (SnowyMountains) Temperature() float64 {
	return 0
}

// Rainfall ...
func (SnowyMountains) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (SnowyMountains) Depth() float64 {
	return 0.45
}

// Scale ...
func (SnowyMountains) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (SnowyMountains) WaterColour() color.RGBA {
	return color.RGBA{R: 0x11, G: 0x56, B: 0xa7, A: 0xa5}
}

// Tags ...
func (SnowyMountains) Tags() []string {
	return []string{"frozen", "ice", "mountain", "overworld", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs", "spawns_snow_foxes", "spawns_white_rabbits"}
}

// String ...
func (SnowyMountains) String() string {
	return "ice_mountains"
}

// EncodeBiome ...
func (SnowyMountains) EncodeBiome() int {
	return 13
}
