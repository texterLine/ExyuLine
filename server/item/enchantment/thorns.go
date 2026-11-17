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

// Thorns is an enchantment that inflicts damage on attackers.
var Thorns thorns

type thorns struct{}

// Name ...
func (thorns) Name() string {
	return "Thorns"
}

// MaxLevel ...
func (thorns) MaxLevel() int {
	return 3
}

// Cost ...
func (thorns) Cost(level int) (int, int) {
	minCost := 10 + 20*(level-1)
	return minCost, minCost + 50
}

// Rarity ...
func (thorns) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// CompatibleWithEnchantment ...
func (thorns) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (thorns) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Armour)
	return ok
}

// ThornsDamageSource is used for damage caused by thorns.
type ThornsDamageSource struct {
	// Owner is the owner of the armour with the thorns enchantment.
	Owner world.Entity
}

func (ThornsDamageSource) ReducedByResistance() bool { return true }
func (ThornsDamageSource) ReducedByArmour() bool     { return false }
func (ThornsDamageSource) Fire() bool                { return false }
func (ThornsDamageSource) IgnoreTotem() bool         { return false }
