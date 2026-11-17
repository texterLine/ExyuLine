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

// FireProtection is an armour enchantment that decreases fire damage.
var FireProtection fireProtection

type fireProtection struct{}

// Name ...
func (fireProtection) Name() string {
	return "Fire Protection"
}

// MaxLevel ...
func (fireProtection) MaxLevel() int {
	return 4
}

// Cost ...
func (fireProtection) Cost(level int) (int, int) {
	minCost := 10 + (level-1)*8
	return minCost, minCost + 8
}

// Rarity ...
func (fireProtection) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// Modifier returns the base protection modifier for the enchantment.
func (fireProtection) Modifier() float64 {
	return 0.08
}

// CompatibleWithEnchantment ...
func (fireProtection) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != BlastProtection && t != ProjectileProtection && t != Protection
}

// CompatibleWithItem ...
func (fireProtection) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Armour)
	return ok
}
