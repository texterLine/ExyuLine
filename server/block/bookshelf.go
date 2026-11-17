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

// Bookshelf is a decorative block that primarily serves to enhance enchanting with an enchanting table.
type Bookshelf struct {
	solid
	bass
}

// BreakInfo ...
func (b Bookshelf) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, alwaysHarvestable, axeEffective, silkTouchDrop(item.NewStack(item.Book{}, 3), item.NewStack(b, 1)))
}

// FlammabilityInfo ...
func (Bookshelf) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(30, 20, true)
}

// FuelInfo ...
func (Bookshelf) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// EncodeItem ...
func (Bookshelf) EncodeItem() (name string, meta int16) {
	return "minecraft:bookshelf", 0
}

// EncodeBlock ...
func (Bookshelf) EncodeBlock() (string, map[string]any) {
	return "minecraft:bookshelf", nil
}
