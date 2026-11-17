// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// GravellyMountainsPlus ...
type GravellyMountainsPlus struct{}

// Temperature ...
func (GravellyMountainsPlus) Temperature() float64 {
	return 0.2
}

// Rainfall ...
func (GravellyMountainsPlus) Rainfall() float64 {
	return 0.3
}

// Depth ...
func (GravellyMountainsPlus) Depth() float64 {
	return 1
}

// Scale ...
func (GravellyMountainsPlus) Scale() float64 {
	return 0.5
}

// WaterColour ...
func (GravellyMountainsPlus) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0e, G: 0x63, B: 0xab, A: 0xa5}
}

// Tags ...
func (GravellyMountainsPlus) Tags() []string {
	return []string{"animal", "extreme_hills", "forest", "monster", "mutated", "overworld", "spawns_cold_variant_farm_animals"}
}

// String ...
func (GravellyMountainsPlus) String() string {
	return "extreme_hills_plus_trees_mutated"
}

// EncodeBiome ...
func (GravellyMountainsPlus) EncodeBiome() int {
	return 162
}
