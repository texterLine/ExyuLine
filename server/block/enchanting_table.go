// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/block/model"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// EnchantingTable is a block that allows players to spend their experience point levels to enchant tools, weapons,
// books, armour, and certain other items.
type EnchantingTable struct {
	transparent
	bassDrum
	sourceWaterDisplacer
}

// Model ...
func (e EnchantingTable) Model() world.BlockModel {
	return model.EnchantingTable{}
}

// BreakInfo ...
func (e EnchantingTable) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(e)).withBlastResistance(1200)
}

// SideClosed ...
func (EnchantingTable) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// LightEmissionLevel ...
func (EnchantingTable) LightEmissionLevel() uint8 {
	return 7
}

// Activate ...
func (EnchantingTable) Activate(pos cube.Pos, _ cube.Face, tx *world.Tx, u item.User, _ *item.UseContext) bool {
	if opener, ok := u.(ContainerOpener); ok {
		opener.OpenBlockContainer(pos, tx)
		return true
	}
	return false
}

// EncodeItem ...
func (EnchantingTable) EncodeItem() (name string, meta int16) {
	return "minecraft:enchanting_table", 0
}

// EncodeBlock ...
func (EnchantingTable) EncodeBlock() (string, map[string]any) {
	return "minecraft:enchanting_table", nil
}

// EncodeNBT is used to encode the block to NBT, so that the enchanting table book will be rendered properly client-side.
// The actual rotation value doesn't need to be set in the NBT, we just need to write the default NBT for the block.
func (e EnchantingTable) EncodeNBT() map[string]any {
	return map[string]any{"id": "EnchantTable"}
}

// DecodeNBT is used to implement world.NBTer.
func (e EnchantingTable) DecodeNBT(map[string]any) any {
	return e
}
