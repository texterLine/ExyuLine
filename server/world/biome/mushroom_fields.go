// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// MushroomFields ...
type MushroomFields struct{}

// Temperature ...
func (MushroomFields) Temperature() float64 {
	return 0.9
}

// Rainfall ...
func (MushroomFields) Rainfall() float64 {
	return 1
}

// Depth ...
func (MushroomFields) Depth() float64 {
	return 0.2
}

// Scale ...
func (MushroomFields) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (MushroomFields) WaterColour() color.RGBA {
	return color.RGBA{R: 0x8a, G: 0x89, B: 0x97, A: 0xa5}
}

// Tags ...
func (MushroomFields) Tags() []string {
	return []string{"mooshroom_island", "overworld", "spawns_without_patrols"}
}

// String ...
func (MushroomFields) String() string {
	return "mushroom_island"
}

// EncodeBiome ...
func (MushroomFields) EncodeBiome() int {
	return 14
}
