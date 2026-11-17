// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"time"
)

// FletchingTable is a block in villages that turn an unemployed villager into a Fletcher.
type FletchingTable struct {
	solid
	bass
}

// BreakInfo ...
func (f FletchingTable) BreakInfo() BreakInfo {
	return newBreakInfo(2.5, alwaysHarvestable, axeEffective, oneOf(f))
}

// FuelInfo ...
func (FletchingTable) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// EncodeItem ...
func (FletchingTable) EncodeItem() (string, int16) {
	return "minecraft:fletching_table", 0
}

// EncodeBlock ...
func (FletchingTable) EncodeBlock() (string, map[string]any) {
	return "minecraft:fletching_table", nil
}
