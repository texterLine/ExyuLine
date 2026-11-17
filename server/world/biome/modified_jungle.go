// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// ModifiedJungle ...
type ModifiedJungle struct{}

// Temperature ...
func (ModifiedJungle) Temperature() float64 {
	return 0.95
}

// Rainfall ...
func (ModifiedJungle) Rainfall() float64 {
	return 0.9
}

// Depth ...
func (ModifiedJungle) Depth() float64 {
	return 0.2
}

// Scale ...
func (ModifiedJungle) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (ModifiedJungle) WaterColour() color.RGBA {
	return color.RGBA{R: 0x1b, G: 0x9e, B: 0xd8, A: 0xa5}
}

// Tags ...
func (ModifiedJungle) Tags() []string {
	return []string{"animal", "jungle", "monster", "mutated", "overworld_generation", "spawns_warm_variant_farm_animals"}
}

// String ...
func (ModifiedJungle) String() string {
	return "jungle_mutated"
}

// EncodeBiome ...
func (ModifiedJungle) EncodeBiome() int {
	return 149
}
