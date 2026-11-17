// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Tuff is an ornamental rock formed from volcanic ash, occurring in underground blobs below Y=16.
type Tuff struct {
	solid
	bassDrum

	// Chiseled specifies if the tuff is chiseled.
	Chiseled bool
}

// BreakInfo ...
func (t Tuff) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(t)).withBlastResistance(6)
}

// EncodeItem ...
func (t Tuff) EncodeItem() (name string, meta int16) {
	if t.Chiseled {
		return "minecraft:chiseled_tuff", 0
	}
	return "minecraft:tuff", 0
}

// EncodeBlock ...
func (t Tuff) EncodeBlock() (string, map[string]any) {
	if t.Chiseled {
		return "minecraft:chiseled_tuff", nil
	}
	return "minecraft:tuff", nil
}
