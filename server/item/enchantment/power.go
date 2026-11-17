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

// Power is a bow enchantment which increases arrow damage.
var Power power

type power struct{}

// Name ...
func (power) Name() string {
	return "Power"
}

// MaxLevel ...
func (power) MaxLevel() int {
	return 5
}

// Cost ...
func (power) Cost(level int) (int, int) {
	minCost := 1 + (level-1)*10
	return minCost, minCost + 15
}

// Rarity ...
func (power) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityCommon
}

// PowerDamage returns the extra base damage dealt by the enchantment and level.
func (power) PowerDamage(level int) float64 {
	return float64(level+1) * 0.5
}

// CompatibleWithEnchantment ...
func (power) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (power) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Bow)
	return ok
}
