// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// GiantSpruceTaigaHills ...
type GiantSpruceTaigaHills struct{}

// Temperature ...
func (GiantSpruceTaigaHills) Temperature() float64 {
	return 0.3
}

// Rainfall ...
func (GiantSpruceTaigaHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (GiantSpruceTaigaHills) Depth() float64 {
	return 0.55
}

// Scale ...
func (GiantSpruceTaigaHills) Scale() float64 {
	return 0.5
}

// WaterColour ...
func (GiantSpruceTaigaHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x28, G: 0x63, B: 0x78, A: 0xa5}
}

// Tags ...
func (GiantSpruceTaigaHills) Tags() []string {
	return []string{"animal", "forest", "hills", "mega", "monster", "mutated", "taiga", "overworld_generation", "spawns_cold_variant_farm_animals"}
}

// String ...
func (GiantSpruceTaigaHills) String() string {
	return "redwood_taiga_hills_mutated"
}

// EncodeBiome ...
func (GiantSpruceTaigaHills) EncodeBiome() int {
	return 161
}
