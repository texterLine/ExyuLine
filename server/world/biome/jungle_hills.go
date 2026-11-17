// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package biome

import "image/color"

// JungleHills ...
type JungleHills struct{}

// Temperature ...
func (JungleHills) Temperature() float64 {
	return 0.95
}

// Rainfall ...
func (JungleHills) Rainfall() float64 {
	return 0.9
}

// Depth ...
func (JungleHills) Depth() float64 {
	return 0.45
}

// Scale ...
func (JungleHills) Scale() float64 {
	return 0.3
}

// WaterColour ...
func (JungleHills) WaterColour() color.RGBA {
	return color.RGBA{R: 0x1b, G: 0x9e, B: 0xd8, A: 0xa5}
}

// Tags ...
func (JungleHills) Tags() []string {
	return []string{"animal", "hills", "jungle", "monster", "overworld", "spawns_jungle_mobs", "spawns_warm_variant_farm_animals"}
}

// String ...
func (JungleHills) String() string {
	return "jungle_hills"
}

// EncodeBiome ...
func (JungleHills) EncodeBiome() int {
	return 22
}
