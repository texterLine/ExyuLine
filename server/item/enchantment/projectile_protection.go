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

// ProjectileProtection is an armour enchantment that reduces damage from
// projectiles.
var ProjectileProtection projectileProtection

type projectileProtection struct{}

// Name ...
func (projectileProtection) Name() string {
	return "Projectile Protection"
}

// MaxLevel ...
func (projectileProtection) MaxLevel() int {
	return 4
}

// Cost ...
func (projectileProtection) Cost(level int) (int, int) {
	minCost := 3 + (level-1)*6
	return minCost, minCost + 6
}

// Rarity ...
func (projectileProtection) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// Modifier returns the base protection modifier for the enchantment.
func (projectileProtection) Modifier() float64 {
	return 0.08
}

// CompatibleWithEnchantment ...
func (projectileProtection) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != BlastProtection && t != FireProtection && t != Protection
}

// CompatibleWithItem ...
func (projectileProtection) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Armour)
	return ok
}
