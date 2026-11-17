// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// BlueIce is a solid block similar to packed ice.
type BlueIce struct {
	solid
}

// BreakInfo ...
func (b BlueIce) BreakInfo() BreakInfo {
	return newBreakInfo(2.8, alwaysHarvestable, pickaxeEffective, silkTouchOnlyDrop(b))
}

// Friction ...
func (b BlueIce) Friction() float64 {
	return 0.989
}

// EncodeItem ...
func (BlueIce) EncodeItem() (name string, meta int16) {
	return "minecraft:blue_ice", 0
}

// EncodeBlock ...
func (BlueIce) EncodeBlock() (string, map[string]any) {
	return "minecraft:blue_ice", nil
}
