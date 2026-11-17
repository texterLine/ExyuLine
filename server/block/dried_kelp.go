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

// DriedKelp is a block primarily used as fuel in furnaces.
type DriedKelp struct {
	solid
}

// BreakInfo ...
func (d DriedKelp) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, hoeEffective, oneOf(d)).withBlastResistance(2.5)
}

// FlammabilityInfo ...
func (DriedKelp) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(30, 5, false)
}

// FuelInfo ...
func (DriedKelp) FuelInfo() item.FuelInfo {
	return newFuelInfo(time.Second * 200)
}

// CompostChance ...
func (DriedKelp) CompostChance() float64 {
	return 0.5
}

// EncodeItem ...
func (DriedKelp) EncodeItem() (name string, meta int16) {
	return "minecraft:dried_kelp_block", 0
}

// EncodeBlock ...
func (DriedKelp) EncodeBlock() (string, map[string]any) {
	return "minecraft:dried_kelp_block", nil
}
