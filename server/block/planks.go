// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// Planks are common blocks used in crafting recipes. They are made by crafting logs into planks.
type Planks struct {
	solid
	bass

	// Wood is the type of wood of the planks. This field must have one of the values found in the material
	// package.
	Wood WoodType
}

// FlammabilityInfo ...
func (p Planks) FlammabilityInfo() FlammabilityInfo {
	if !p.Wood.Flammable() {
		return newFlammabilityInfo(0, 0, false)
	}
	return newFlammabilityInfo(5, 20, true)
}

// BreakInfo ...
func (p Planks) BreakInfo() BreakInfo {
	return newBreakInfo(2, alwaysHarvestable, axeEffective, oneOf(p)).withBlastResistance(3)
}

// RepairsWoodTools ...
func (p Planks) RepairsWoodTools() bool {
	return true
}

// FuelInfo ...
func (p Planks) FuelInfo() item.FuelInfo {
	if !p.Wood.Flammable() {
		return item.FuelInfo{}
	}
	return newFuelInfo(time.Second * 15)
}

// EncodeItem ...
func (p Planks) EncodeItem() (name string, meta int16) {
	return "minecraft:" + p.Wood.String() + "_planks", 0
}

// EncodeBlock ...
func (p Planks) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + p.Wood.String() + "_planks", nil
}

// allPlanks returns all planks types.
func allPlanks() (planks []world.Block) {
	for _, w := range WoodTypes() {
		planks = append(planks, Planks{Wood: w})
	}
	return
}
