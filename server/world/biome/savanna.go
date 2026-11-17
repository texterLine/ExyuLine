// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Savanna ...
type Savanna struct{}

// Temperature ...
func (Savanna) Temperature() float64 {
	return 1.2
}

// Rainfall ...
func (Savanna) Rainfall() float64 {
	return 0
}

// Depth ...
func (Savanna) Depth() float64 {
	return 0.125
}

// Scale ...
func (Savanna) Scale() float64 {
	return 0.05
}

// WaterColour ...
func (Savanna) WaterColour() color.RGBA {
	return color.RGBA{R: 0x2c, G: 0x8b, B: 0x9c, A: 0xa5}
}

// Tags ...
func (Savanna) Tags() []string {
	return []string{"animal", "monster", "overworld", "savanna", "spawns_savanna_mobs", "spawns_warm_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (Savanna) String() string {
	return "savanna"
}

// EncodeBiome ...
func (Savanna) EncodeBiome() int {
	return 35
}
