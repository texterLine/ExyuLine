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

// SwiftSneak is a non-renewable enchantment that can be applied to leggings
// and allows the player to walk more quickly while sneaking.
var SwiftSneak swiftSneak

type swiftSneak struct{}

// Name ...
func (swiftSneak) Name() string {
	return "Swift Sneak"
}

// MaxLevel ...
func (swiftSneak) MaxLevel() int {
	return 3
}

// Cost ...
func (swiftSneak) Cost(level int) (int, int) {
	minCost := level * 25
	return minCost, minCost + 50
}

// Rarity ...
func (swiftSneak) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// CompatibleWithEnchantment ...
func (swiftSneak) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// Treasure ...
func (swiftSneak) Treasure() bool {
	return true
}

// CompatibleWithItem ...
func (swiftSneak) CompatibleWithItem(i world.Item) bool {
	b, ok := i.(item.LeggingsType)
	return ok && b.Leggings()
}
