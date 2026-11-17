// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Badlands ...
type Badlands struct{}

// Temperature ...
func (Badlands) Temperature() float64 {
	return 2
}

// Rainfall ...
func (Badlands) Rainfall() float64 {
	return 0
}

// Depth ...
func (Badlands) Depth() float64 {
	return 0.1
}

// Scale ...
func (Badlands) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (Badlands) WaterColour() color.RGBA {
	return color.RGBA{R: 0x4e, G: 0x7f, B: 0x81, A: 0xa5}
}

// Tags ...
func (Badlands) Tags() []string {
	return []string{"animal", "mesa", "monster", "overworld", "spawns_mesa_mobs", "spawns_warm_variant_farm_animals", "surface_mineshaft"}
}

// String ...
func (Badlands) String() string {
	return "mesa"
}

// EncodeBiome ...
func (Badlands) EncodeBiome() int {
	return 37
}
