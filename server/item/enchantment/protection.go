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

// Protection is an armour enchantment which increases the damage reduction.
var Protection protection

type protection struct{}

// Name ...
func (protection) Name() string {
	return "Protection"
}

// MaxLevel ...
func (protection) MaxLevel() int {
	return 4
}

// Cost ...
func (protection) Cost(level int) (int, int) {
	minCost := 1 + (level-1)*11
	return minCost, minCost + 11
}

// Rarity ...
func (protection) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityCommon
}

// Modifier returns the base protection modifier for the enchantment.
func (protection) Modifier() float64 {
	return 0.04
}

// CompatibleWithEnchantment ...
func (protection) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != BlastProtection && t != FireProtection && t != ProjectileProtection
}

// CompatibleWithItem ...
func (protection) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Armour)
	return ok
}
