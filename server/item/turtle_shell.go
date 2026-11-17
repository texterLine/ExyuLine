// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// TurtleShell are items that are used for brewing or as a helmet to give the player the Water Breathing
// status effect.
type TurtleShell struct{}

// Use handles the using of a turtle shell to auto-equip it in an armour slot.
func (TurtleShell) Use(_ *world.Tx, _ User, ctx *UseContext) bool {
	ctx.SwapHeldWithArmour(0)
	return false
}

// DurabilityInfo ...
func (TurtleShell) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: 276,
		BrokenItem:    simpleItem(Stack{}),
	}
}

// RepairableBy ...
func (TurtleShell) RepairableBy(i Stack) bool {
	_, ok := i.Item().(Scute)
	return ok
}

// MaxCount always returns 1.
func (TurtleShell) MaxCount() int {
	return 1
}

// DefencePoints ...
func (TurtleShell) DefencePoints() float64 {
	return 2
}

// Toughness ...
func (TurtleShell) Toughness() float64 {
	return 0
}

// KnockBackResistance ...
func (TurtleShell) KnockBackResistance() float64 {
	return 0
}

// EnchantmentValue ...
func (TurtleShell) EnchantmentValue() int {
	return 9
}

// Helmet ...
func (TurtleShell) Helmet() bool {
	return true
}

// EncodeItem ...
func (TurtleShell) EncodeItem() (name string, meta int16) {
	return "minecraft:turtle_helmet", 0
}
