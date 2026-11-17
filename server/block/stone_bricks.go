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
)

// StoneBricks are materials found in structures such as strongholds, igloo basements, jungle temples, ocean ruins
// and ruined portals.
type StoneBricks struct {
	solid
	bassDrum

	// Type is the type of stone bricks of the block.
	Type StoneBricksType
}

// BreakInfo ...
func (s StoneBricks) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(s)).withBlastResistance(6)
}

// SmeltInfo ...
func (s StoneBricks) SmeltInfo() item.SmeltInfo {
	if s.Type == NormalStoneBricks() {
		return newSmeltInfo(item.NewStack(StoneBricks{Type: CrackedStoneBricks()}, 1), 0.1)
	}
	return item.SmeltInfo{}
}

// EncodeItem ...
func (s StoneBricks) EncodeItem() (name string, meta int16) {
	return "minecraft:" + s.Type.String(), 0
}

// EncodeBlock ...
func (s StoneBricks) EncodeBlock() (string, map[string]any) {
	return "minecraft:" + s.Type.String(), nil
}

// allStoneBricks returns a list of all stoneBricks block variants.
func allStoneBricks() (s []world.Block) {
	for _, t := range StoneBricksTypes() {
		s = append(s, StoneBricks{Type: t})
	}
	return
}
