// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// Forest ...
type Forest struct{}

// Temperature ...
func (Forest) Temperature() float64 {
	return 0.7
}

// Rainfall ...
func (Forest) Rainfall() float64 {
	return 0.8
}

// Depth ...
func (Forest) Depth() float64 {
	return 0.1
}

// Scale ...
func (Forest) Scale() float64 {
	return 0.2
}

// WaterColour ...
func (Forest) WaterColour() color.RGBA {
	return color.RGBA{R: 0x1e, G: 0x97, B: 0xf2, A: 0xa5}
}

// Tags ...
func (Forest) Tags() []string {
	return []string{"animal", "forest", "monster", "overworld", "bee_habitat", "has_structure_abandoned_camp", "oak_forest"}
}

// String ...
func (Forest) String() string {
	return "forest"
}

// EncodeBiome ...
func (Forest) EncodeBiome() int {
	return 4
}
