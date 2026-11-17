// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Desert ...
type Desert struct{}

// Temperature ...
func (Desert) Temperature() float64 {
	return 2
}

// Rainfall ...
func (Desert) Rainfall() float64 {
	return 0
}

// Depth ...
func (Desert) Depth() float64 {
	return 0.125
}

// Scale ...
func (Desert) Scale() float64 {
	return 0.05
}

// WaterColour ...
func (Desert) WaterColour() color.RGBA {
	return color.RGBA{R: 0x32, G: 0xa5, B: 0x98, A: 0xa5}
}

// Tags ...
func (Desert) Tags() []string {
	return []string{"desert", "monster", "overworld", "spawns_gold_rabbits", "spawns_warm_variant_farm_animals", "spawns_warm_variant_frogs"}
}

// String ...
func (Desert) String() string {
	return "desert"
}

// EncodeBiome ...
func (Desert) EncodeBiome() int {
	return 2
}
