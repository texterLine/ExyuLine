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

// Iron is a precious metal block made from 9 iron ingots.
type Iron struct {
	solid
}

// Instrument ...
func (i Iron) Instrument() sound.Instrument {
	return sound.IronXylophone()
}

// BreakInfo ...
func (i Iron) BreakInfo() BreakInfo {
	return newBreakInfo(5, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, pickaxeEffective, oneOf(i)).withBlastResistance(6)
}

// PowersBeacon ...
func (Iron) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (Iron) EncodeItem() (name string, meta int16) {
	return "minecraft:iron_block", 0
}

// EncodeBlock ...
func (Iron) EncodeBlock() (string, map[string]any) {
	return "minecraft:iron_block", nil
}
