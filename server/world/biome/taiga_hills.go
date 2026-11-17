// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// TaigaHills ...
type TaigaHills struct{}

// Temperature ...
func (TaigaHills) Temperature() float64 {
	return 0.25
}

// Rainfall ...
func (TaigaHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (TaigaHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (TaigaHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (TaigaHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x23, G: 0x65, B: 0x83, A: 0xa5}
}

// Tags ...
func (TaigaHills) Tags() []string {
	return []string{"animal", "hills", "monster", "overworld", "forest", "taiga", "spawns_cold_variant_farm_animals"}
}

// String ...
func (TaigaHills) String() string {
	return "taiga_hills"
}

// EncodeBiome ...
func (TaigaHills) EncodeBiome() int {
	return 19
}
