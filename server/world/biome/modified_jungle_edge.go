// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// ModifiedJungleEdge ...
type ModifiedJungleEdge struct{}

// Temperature ...
func (ModifiedJungleEdge) Temperature() float64 {
	return 0.95
}

// Rainfall ...
func (ModifiedJungleEdge) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (ModifiedJungleEdge) Depth() float64 {
	return 0.2
}

// Scale ...
func (ModifiedJungleEdge) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (ModifiedJungleEdge) WaterColour() color.RGBA {
	return color.RGBA{R: 0x0d, G: 0x8a, B: 0xe3, A: 0xa5}
}

// Tags ...
func (ModifiedJungleEdge) Tags() []string {
	return []string{"animal", "edge", "jungle", "monster", "mutated", "overworld_generation", "spawns_jungle_mobs", "spawns_warm_variant_farm_animals"}
}

// String ...
func (ModifiedJungleEdge) String() string {
	return "jungle_edge_mutated"
}

// EncodeBiome ...
func (ModifiedJungleEdge) EncodeBiome() int {
	return 151
}
