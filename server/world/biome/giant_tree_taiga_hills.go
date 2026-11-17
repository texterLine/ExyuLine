// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// GiantTreeTaigaHills ...
type GiantTreeTaigaHills struct{}

// Temperature ...
func (GiantTreeTaigaHills) Temperature() float64 {
	return 0.3
}

// Rainfall ...
func (GiantTreeTaigaHills) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (GiantTreeTaigaHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (GiantTreeTaigaHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (GiantTreeTaigaHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x28, G: 0x63, B: 0x78, A: 0xa5}
}

// Tags ...
func (GiantTreeTaigaHills) Tags() []string {
	return []string{"animal", "forest", "hills", "mega", "monster", "overworld", "taiga", "spawns_cold_variant_farm_animals"}
}

// String ...
func (GiantTreeTaigaHills) String() string {
	return "mega_taiga_hills"
}

// EncodeBiome ...
func (GiantTreeTaigaHills) EncodeBiome() int {
	return 33
}
