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

// BambooMosaic is a decorative bamboo plank variant.
type BambooMosaic struct {
	solid
	bass
}

// FlammabilityInfo ...
func (BambooMosaic) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(5, 20, true)
}

// BreakInfo ...
func (b BambooMosaic) BreakInfo() BreakInfo {
	return newBreakInfo(2, alwaysHarvestable, axeEffective, oneOf(b)).withBlastResistance(3)
}

// RepairsWoodTools ...
func (BambooMosaic) RepairsWoodTools() bool {
	return true
}

// FuelInfo ...
func (BambooMosaic) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 15)
}

// EncodeItem ...
func (BambooMosaic) EncodeItem() (name string, meta int16) {
	return "minecraft:bamboo_mosaic", 0
}

// EncodeBlock ...
func (BambooMosaic) EncodeBlock() (string, map[string]any) {
	return "minecraft:bamboo_mosaic", nil
}
