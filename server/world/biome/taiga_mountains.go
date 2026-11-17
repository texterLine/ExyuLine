// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// TaigaMountains ...
type TaigaMountains struct{}

// Temperature ...
func (TaigaMountains) Temperature() float64 {
	return 0.25
}

// Rainfall ...
func (TaigaMountains) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (TaigaMountains) Depth() float64 {
	return 0.2
}

// Scale ...
func (TaigaMountains) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (TaigaMountains) WaterColour() color.RGBA {
	return color.RGBA{R: 0x1e, G: 0x6b, B: 0x82, A: 0xa5}
}

// Tags ...
func (TaigaMountains) Tags() []string {
	return []string{"animal", "forest", "monster", "mutated", "taiga", "overworld_generation", "spawns_cold_variant_farm_animals"}
}

// String ...
func (TaigaMountains) String() string {
	return "taiga_mutated"
}

// EncodeBiome ...
func (TaigaMountains) EncodeBiome() int {
	return 133
}
