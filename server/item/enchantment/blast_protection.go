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

// BlastProtection is an armour enchantment that reduces damage from explosions.
var BlastProtection blastProtection

type blastProtection struct{}

// Name ...
func (blastProtection) Name() string {
	return "Blast Protection"
}

// MaxLevel ...
func (blastProtection) MaxLevel() int {
	return 4
}

// Cost ...
func (blastProtection) Cost(level int) (int, int) {
	minCost := 5 + (level-1)*8
	return minCost, minCost + 8
}

// Rarity ...
func (blastProtection) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// Modifier returns the base protection modifier for the enchantment.
func (blastProtection) Modifier() float64 {
	return 0.08
}

// CompatibleWithEnchantment ...
func (blastProtection) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != FireProtection && t != ProjectileProtection && t != Protection
}

// CompatibleWithItem ...
func (blastProtection) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Armour)
	return ok
}
