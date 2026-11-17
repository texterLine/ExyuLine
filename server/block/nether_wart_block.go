// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// NetherWartBlock is a decorative block found in crimson forests and crafted using Nether wart.
type NetherWartBlock struct {
	solid

	// Warped is the turquoise variant found in warped forests, but cannot be crafted unlike Nether wart block.
	Warped bool
}

// BreakInfo ...
func (n NetherWartBlock) BreakInfo() BreakInfo {
	return newBreakInfo(1, alwaysHarvestable, hoeEffective, oneOf(n))
}

// CompostChance ...
func (NetherWartBlock) CompostChance() float64 {
	return 0.85
}

// EncodeItem ...
func (n NetherWartBlock) EncodeItem() (name string, meta int16) {
	if n.Warped {
		return "minecraft:warped_wart_block", 0
	}
	return "minecraft:nether_wart_block", 0
}

// EncodeBlock ...
func (n NetherWartBlock) EncodeBlock() (name string, properties map[string]interface{}) {
	if n.Warped {
		return "minecraft:warped_wart_block", nil
	}
	return "minecraft:nether_wart_block", nil
}
