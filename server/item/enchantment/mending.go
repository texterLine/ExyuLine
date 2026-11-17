// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package enchantment

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// Mending is an enchantment that repairs the item when experience orbs are
// collected.
var Mending mending

type mending struct{}

// Name ...
func (mending) Name() string {
	return "Mending"
}

// MaxLevel ...
func (mending) MaxLevel() int {
	return 1
}

// Cost ...
func (mending) Cost(level int) (int, int) {
	minCost := level * 25
	return minCost, minCost + 50
}

// Rarity ...
func (mending) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// Treasure ...
func (mending) Treasure() bool {
	return true
}

// CompatibleWithEnchantment ...
func (mending) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != Infinity
}

// CompatibleWithItem ...
func (mending) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Durable)
	return ok
}
