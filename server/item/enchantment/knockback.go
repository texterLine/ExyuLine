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

// Knockback is an enchantment to a sword that increases the sword's knock-back.
var Knockback knockback

type knockback struct{}

// Name ...
func (knockback) Name() string {
	return "Knockback"
}

// MaxLevel ...
func (knockback) MaxLevel() int {
	return 2
}

// Cost ...
func (knockback) Cost(level int) (int, int) {
	minCost := 5 + (level-1)*20
	return minCost, minCost + 50
}

// Rarity ...
func (knockback) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// Force returns the increase in knock-back force from the enchantment.
func (knockback) Force(level int) float64 {
	return float64(level) / 2
}

// CompatibleWithEnchantment ...
func (knockback) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (knockback) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && t.ToolType() == item.TypeSword
}
