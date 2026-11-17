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

// DepthStrider is a boot enchantment that increases underwater movement speed.
var DepthStrider depthStrider

type depthStrider struct{}

// Name ...
func (depthStrider) Name() string {
	return "Depth Strider"
}

// MaxLevel ...
func (depthStrider) MaxLevel() int {
	return 3
}

// Cost ...
func (depthStrider) Cost(level int) (int, int) {
	minCost := level * 10
	return minCost, minCost + 15
}

// Rarity ...
func (depthStrider) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (depthStrider) CompatibleWithEnchantment(item.EnchantmentType) bool {
	// TODO: Frost Walker
	return true
}

// CompatibleWithItem ...
func (depthStrider) CompatibleWithItem(i world.Item) bool {
	b, ok := i.(item.BootsType)
	return ok && b.Boots()
}
