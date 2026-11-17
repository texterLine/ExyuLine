// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// FrozenRiver ...
type FrozenRiver struct{}

// Temperature ...
func (FrozenRiver) Temperature() float64 {
	return 0
}

// Rainfall ...
func (FrozenRiver) Rainfall() float64 {
	return 0.5
}

// Depth ...
func (FrozenRiver) Depth() float64 {
	return -0.5
}

// Scale ...
func (FrozenRiver) Scale() float64 {
	return 0
}

// WaterColour ...
func (FrozenRiver) WaterColour() color.RGBA {
	return color.RGBA{R: 0x18, G: 0x53, B: 0x90, A: 0xa5}
}

// Tags ...
func (FrozenRiver) Tags() []string {
	return []string{"frozen", "overworld", "river", "spawns_cold_variant_farm_animals", "spawns_cold_variant_frogs", "spawns_river_mobs", "spawns_snow_foxes", "spawns_white_rabbits"}
}

// String ...
func (FrozenRiver) String() string {
	return "frozen_river"
}

// EncodeBiome ...
func (FrozenRiver) EncodeBiome() int {
	return 11
}
