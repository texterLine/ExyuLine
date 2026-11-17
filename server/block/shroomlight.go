// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Shroomlight are light-emitting blocks that generate in huge fungi.
type Shroomlight struct {
	solid
}

// LightEmissionLevel ...
func (Shroomlight) LightEmissionLevel() uint8 {
	return 15
}

// BreakInfo ...
func (s Shroomlight) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, hoeEffective, oneOf(s))
}

// CompostChance ...
func (Shroomlight) CompostChance() float64 {
	return 0.65
}

// EncodeItem ...
func (Shroomlight) EncodeItem() (name string, meta int16) {
	return "minecraft:shroomlight", 0
}

// EncodeBlock ...
func (Shroomlight) EncodeBlock() (string, map[string]any) {
	return "minecraft:shroomlight", nil
}
