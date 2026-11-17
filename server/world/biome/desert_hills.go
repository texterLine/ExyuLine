// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// DesertHills ...
type DesertHills struct{}

// Temperature ...
func (DesertHills) Temperature() float64 {
	return 2
}

// Rainfall ...
func (DesertHills) Rainfall() float64 {
	return 0
}

// Depth ...
func (DesertHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (DesertHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (DesertHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x1a, G: 0x7a, B: 0xa1, A: 0xa5}
}

// Tags ...
func (DesertHills) Tags() []string {
	return []string{"desert", "hills", "monster", "overworld", "spawns_gold_rabbits", "spawns_warm_variant_farm_animals"}
}

// String ...
func (DesertHills) String() string {
	return "desert_hills"
}

// EncodeBiome ...
func (DesertHills) EncodeBiome() int {
	return 17
}
