// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DesertLakes ...
type DesertLakes struct{}

// Temperature ...
func (DesertLakes) Temperature() float64 {
	return 2
}

// Rainfall ...
func (DesertLakes) Rainfall() float64 {
	return 0
}

// Depth ...
func (DesertLakes) Depth() float64 {
	return 0.225
}

// Scale ...
func (DesertLakes) Scale() float64 {
	return 0.25
}

// WaterColour ...
func (DesertLakes) WaterColour() color.RGBA {
	return color.RGBA{R: 0x32, G: 0xa5, B: 0x98, A: 0xa5}
}

// Tags ...
func (DesertLakes) Tags() []string {
	return []string{"desert", "monster", "mutated", "overworld_generation", "spawns_gold_rabbits", "spawns_warm_variant_farm_animals"}
}

// String ...
func (DesertLakes) String() string {
	return "desert_mutated"
}

// EncodeBiome ...
func (DesertLakes) EncodeBiome() int {
	return 130
}
