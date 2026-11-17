// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// BambooJungle ...
type BambooJungle struct{}

// Temperature ...
func (BambooJungle) Temperature() float64 {
	return 0.95
}

// Rainfall ...
func (BambooJungle) Rainfall() float64 {
	return 0.9
}

// Depth ...
func (BambooJungle) Depth() float64 {
	return 0.1
}

// Scale ...
func (BambooJungle) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (BambooJungle) WaterColour() color.RGBA {
	return color.RGBA{R: 0x14, G: 0xa2, B: 0xc5, A: 0xa5}
}

// Tags ...
func (BambooJungle) Tags() []string {
	return []string{"animal", "bamboo", "jungle", "monster", "overworld", "spawns_jungle_mobs", "spawns_warm_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (BambooJungle) String() string {
	return "bamboo_jungle"
}

// EncodeBiome ...
func (BambooJungle) EncodeBiome() int {
	return 48
}
