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

// SoulSpeed is an enchantment that can be applied on boots and allows the
// player to walk more quickly on soul sand or soul soil.
var SoulSpeed soulSpeed

type soulSpeed struct{}

// Name ...
func (soulSpeed) Name() string {
	return "Soul Speed"
}

// MaxLevel ...
func (soulSpeed) MaxLevel() int {
	return 3
}

// Cost ...
func (soulSpeed) Cost(level int) (int, int) {
	minCost := level * 10
	return minCost, minCost + 15
}

// Rarity ...
func (soulSpeed) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// Treasure ...
func (soulSpeed) Treasure() bool {
	return true
}

// CompatibleWithEnchantment ...
func (soulSpeed) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (soulSpeed) CompatibleWithItem(i world.Item) bool {
	b, ok := i.(item.BootsType)
	return ok && b.Boots()
}
