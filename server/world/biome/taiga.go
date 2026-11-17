// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Taiga ...
type Taiga struct{}

// Temperature ...
func (Taiga) Temperature() float64 {
	return 0.25
}

// Rainfall ...
func (Taiga) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (Taiga) Depth() float64 {
	return 0.1
}

// Scale ...
func (Taiga) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (Taiga) WaterColour() color.RGBA {
	return color.RGBA{R: 0x28, G: 0x70, B: 0x82, A: 0xa5}
}

// Tags ...
func (Taiga) Tags() []string {
	return []string{"animal", "forest", "monster", "overworld", "taiga", "has_structure_trail_ruins", "spawns_cold_variant_farm_animals", "has_structure_abandoned_camp"}
}

// String ...
func (Taiga) String() string {
	return "taiga"
}

// EncodeBiome ...
func (Taiga) EncodeBiome() int {
	return 5
}
