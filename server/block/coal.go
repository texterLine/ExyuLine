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

// Coal is a precious mineral block made from 9 coal.
type Coal struct {
	solid
	bassDrum
}

// BreakInfo ...
func (c Coal) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// FlammabilityInfo ...
func (Coal) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(5, 5, false)
}

// FuelInfo ...
func (Coal) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 800)
}

// EncodeItem ...
func (Coal) EncodeItem() (name string, meta int16) {
	return "minecraft:coal_block", 0
}

// EncodeBlock ...
func (Coal) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:coal_block", nil
}
