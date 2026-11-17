// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// PolishedBlackstoneBrick are a brick variant of polished blackstone and can spawn in bastion remnants and ruined portals.
type PolishedBlackstoneBrick struct {
	solid
	bassDrum

	// Cracked specifies if the polished blackstone bricks is its cracked variant.
	Cracked bool
}

// BreakInfo ...
func (b PolishedBlackstoneBrick) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(b)).withBlastResistance(6)
}

// SmeltInfo ...
func (b PolishedBlackstoneBrick) SmeltInfo() item.SmeltInfo {
	if b.Cracked {
		return item.SmeltInfo{}
	}
	return newSmeltInfo(item.NewStack(PolishedBlackstoneBrick{Cracked: true}, 1), 0.1)
}

// EncodeItem ...
func (b PolishedBlackstoneBrick) EncodeItem() (name string, meta int16) {
	name = "polished_blackstone_bricks"
	if b.Cracked {
		name = "cracked_" + name
	}
	return "minecraft:" + name, 0
}

// EncodeBlock ...
func (b PolishedBlackstoneBrick) EncodeBlock() (string, map[string]any) {
	name := "polished_blackstone_bricks"
	if b.Cracked {
		name = "cracked_" + name
	}
	return "minecraft:" + name, nil
}
