// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Cinnabar is a decorative rock that generates throughout sulfur caves and as part of sulfur springs.
type Cinnabar struct {
	solid
	bassDrum

	// Chiseled specifies if the cinnabar is chiseled.
	Chiseled bool
}

// BreakInfo ...
func (c Cinnabar) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// EncodeItem ...
func (c Cinnabar) EncodeItem() (name string, meta int16) {
	if c.Chiseled {
		return "minecraft:chiseled_cinnabar", 0
	}
	return "minecraft:cinnabar", 0
}

// EncodeBlock ...
func (c Cinnabar) EncodeBlock() (string, map[string]any) {
	if c.Chiseled {
		return "minecraft:chiseled_cinnabar", nil
	}
	return "minecraft:cinnabar", nil
}
