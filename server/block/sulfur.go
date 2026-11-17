// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Sulfur is a decorative rock that generates throughout sulfur caves and as part of sulfur springs.
type Sulfur struct {
	solid
	bassDrum

	// Chiseled specifies if the sulfur is chiseled.
	Chiseled bool
}

// BreakInfo ...
func (s Sulfur) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(s)).withBlastResistance(6)
}

// EncodeItem ...
func (s Sulfur) EncodeItem() (name string, meta int16) {
	if s.Chiseled {
		return "minecraft:chiseled_sulfur", 0
	}
	return "minecraft:sulfur", 0
}

// EncodeBlock ...
func (s Sulfur) EncodeBlock() (string, map[string]any) {
	if s.Chiseled {
		return "minecraft:chiseled_sulfur", nil
	}
	return "minecraft:sulfur", nil
}
