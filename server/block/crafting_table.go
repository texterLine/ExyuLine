// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// CraftingTable is a utility block that allows the player to craft a variety of blocks and items.
type CraftingTable struct {
	bass
	solid
}

// EncodeItem ...
func (c CraftingTable) EncodeItem() (name string, meta int16) {
	return "minecraft:crafting_table", 0
}

// EncodeBlock ...
func (c CraftingTable) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:crafting_table", nil
}

// BreakInfo ...
func (c CraftingTable) BreakInfo() BreakInfo {
	return newBreakInfo(2.5, alwaysHarvestable, axeEffective, oneOf(c))
}

// FuelInfo ...
func (CraftingTable) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// Activate ...
func (c CraftingTable) Activate(pos cube.Pos, _ cube.Face, tx *world.Tx, u item.User, _ *item.UseContext) bool {
	if opener, ok := u.(ContainerOpener); ok {
		opener.OpenBlockContainer(pos, tx)
		return true
	}
	return false
}
