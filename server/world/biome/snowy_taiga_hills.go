// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// SnowyTaigaHills ...
type SnowyTaigaHills struct{}

// Temperature ...
func (SnowyTaigaHills) Temperature() float64 {
	return -0.5
}

// Rainfall ...
func (SnowyTaigaHills) Rainfall() float64 {
	return 0.4
}

// Depth ...
func (SnowyTaigaHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (SnowyTaigaHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (SnowyTaigaHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x24, G: 0x5b, B: 0x78, A: 0xa5}
}

// Tags ...
func (SnowyTaigaHills) Tags() []string {
	return []string{"animal", "cold", "forest", "hills", "monster", "overworld", "taiga", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs", "spawns_white_rabbits"}
}

// String ...
func (SnowyTaigaHills) String() string {
	return "cold_taiga_hills"
}

// EncodeBiome ...
func (SnowyTaigaHills) EncodeBiome() int {
	return 31
}
