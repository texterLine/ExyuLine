// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// MountainEdge ...
type MountainEdge struct{}

// Temperature ...
func (MountainEdge) Temperature() float64 {
	return 0.2
}

// Rainfall ...
func (MountainEdge) Rainfall() float64 {
	return 0.3
}

// Depth ...
func (MountainEdge) Depth() float64 {
	return 0.8
}

// Scale ...
func (MountainEdge) Scale() float64 {
	return 0.4
}

// WaterColour ...
func (MountainEdge) WaterColour() color.RGBA {
	return color.RGBA{R: 0x04, G: 0x5c, B: 0xd5, A: 0xa5}
}

// Tags ...
func (MountainEdge) Tags() []string {
	return []string{"animal", "edge", "extreme_hills", "monster", "mountain", "overworld", "spawns_cold_variant_farm_animals"}
}

// String ...
func (MountainEdge) String() string {
	return "extreme_hills_edge"
}

// EncodeBiome ...
func (MountainEdge) EncodeBiome() int {
	return 20
}
