// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Swamp ...
type Swamp struct{}

// Temperature ...
func (Swamp) Temperature() float64 {
	return 0.8
}

// Rainfall ...
func (Swamp) Rainfall() float64 {
	return 0.9
}

// Depth ...
func (Swamp) Depth() float64 {
	return -0.2
}

// Scale ...
func (Swamp) Scale() float64 {
	return 0.1
}

// WaterColour ...
func (Swamp) WaterColour() color.RGBA {
	return color.RGBA{R: 0x61, G: 0x7b, B: 0x64, A: 0xa5}
}

// Tags ...
func (Swamp) Tags() []string {
	return []string{"animal", "monster", "overworld", "swamp", "spawns_slimes_on_surface", "slime", "swamp_water_huge_mushroom", "has_structure_abandoned_camp"}
}

// String ...
func (Swamp) String() string {
	return "swampland"
}

// EncodeBiome ...
func (Swamp) EncodeBiome() int {
	return 6
}
