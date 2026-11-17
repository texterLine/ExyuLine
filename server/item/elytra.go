// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// Elytra is a pair of rare wings found in end ships that are the only single-item source of flight in Survival mode.
type Elytra struct{}

// Use handles the using of an elytra to auto-equip it in an armour slot.
func (Elytra) Use(_ *world.Tx, _ User, ctx *UseContext) bool {
	ctx.SwapHeldWithArmour(1)
	return false
}

// DurabilityInfo ...
func (Elytra) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: 433,
		Persistent:    true,
		BrokenItem:    simpleItem(Stack{}),
	}
}

// RepairableBy ...
func (Elytra) RepairableBy(i Stack) bool {
	_, ok := i.Item().(PhantomMembrane)
	return ok
}

// MaxCount always returns 1.
func (Elytra) MaxCount() int {
	return 1
}

// Chestplate ...
func (Elytra) Chestplate() bool {
	return true
}

// DefencePoints ...
func (Elytra) DefencePoints() float64 {
	return 0
}

// Toughness ...
func (e Elytra) Toughness() float64 {
	return 0
}

// KnockBackResistance ...
func (e Elytra) KnockBackResistance() float64 {
	return 0
}

// EncodeItem ...
func (Elytra) EncodeItem() (name string, meta int16) {
	return "minecraft:elytra", 0
}
