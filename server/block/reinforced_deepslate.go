// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// ReinforcedDeepslate is a tough decorative block that spawns in ancient cities.
type ReinforcedDeepslate struct {
	solid
	bassDrum
}

// BreakInfo ...
func (r ReinforcedDeepslate) BreakInfo() BreakInfo {
	return newBreakInfo(55, alwaysHarvestable, nothingEffective, oneOf(r)).withBlastResistance(1200)
}

// EncodeItem ...
func (ReinforcedDeepslate) EncodeItem() (name string, meta int16) {
	return "minecraft:reinforced_deepslate", 0
}

// EncodeBlock ...
func (ReinforcedDeepslate) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:reinforced_deepslate", nil
}
