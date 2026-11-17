// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// MangroveSwamp ...
type MangroveSwamp struct{}

// Temperature ...
func (MangroveSwamp) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (MangroveSwamp) Rainfall() float64 {
	return 0.9
}

// Depth ...
func (MangroveSwamp) Depth() float64 {
	return -0.2
}

// Scale ...
func (MangroveSwamp) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (MangroveSwamp) WaterColour() color.RGBA {
	return color.RGBA{R: 0x60, G: 0xb7, B: 0xff, A: 0xa6}
}

// Tags ...
func (MangroveSwamp) Tags() []string {
	return []string{"mangrove_swamp", "overworld", "monster", "spawns_slimes_on_surface", "spawns_warm_variant_farm_animals", "spawns_warm_variant_frogs", "slime"}
}

// String ...
func (MangroveSwamp) String() string {
	return "mangrove_swamp"
}

// EncodeBiome ...
func (MangroveSwamp) EncodeBiome() int {
	return 191
}
