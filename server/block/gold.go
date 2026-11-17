// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world/sound"
)

// Gold is a precious metal block crafted from 9 gold ingots.
type Gold struct {
	solid
}

// Instrument ...
func (g Gold) Instrument() sound.Instrument {
	return sound.Bell()
}

// BreakInfo ...
func (g Gold) BreakInfo() BreakInfo {
	return newBreakInfo(3, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierIron.HarvestLevel
	}, pickaxeEffective, oneOf(g)).withBlastResistance(6)
}

// PowersBeacon ...
func (Gold) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (Gold) EncodeItem() (name string, meta int16) {
	return "minecraft:gold_block", 0
}

// EncodeBlock ...
func (Gold) EncodeBlock() (string, map[string]any) {
	return "minecraft:gold_block", nil
}
