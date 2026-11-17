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
	"math/rand/v2"
)

// Unbreaking is an enchantment that gives a chance for an item to avoid
// durability reduction when it is used, effectively increasing the item's
// durability.
var Unbreaking unbreaking

type unbreaking struct{}

// Name ...
func (unbreaking) Name() string {
	return "Unbreaking"
}

// MaxLevel ...
func (unbreaking) MaxLevel() int {
	return 3
}

// Cost ...
func (unbreaking) Cost(level int) (int, int) {
	minCost := 5 + 8*(level-1)
	return minCost, minCost + 50
}

// Rarity ...
func (unbreaking) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// CompatibleWithEnchantment ...
func (unbreaking) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (unbreaking) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Durable)
	return ok
}

// Reduce returns the amount of damage that should be reduced with unbreaking.
func (unbreaking) Reduce(it world.Item, level, amount int) int {
	after := amount
	_, ok := it.(item.Armour)
	for i := 0; i < amount; i++ {
		if (!ok || rand.Float64() >= 0.6) && rand.IntN(level+1) > 0 {
			after--
		}
	}
	return after
}
