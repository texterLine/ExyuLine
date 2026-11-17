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

// Fortune is an enchantment that gives a chance to receive more item drops from certain blocks.
var Fortune fortune

type fortune struct{}

// Name ...
func (fortune) Name() string {
	return "Fortune"
}

// MaxLevel ...
func (fortune) MaxLevel() int {
	return 3
}

// Cost ...
func (fortune) Cost(level int) (int, int) {
	minCost := 15 + (level-1)*9
	return minCost, minCost + 50 + level
}

// Rarity ...
func (fortune) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (fortune) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != SilkTouch
}

// CompatibleWithItem ...
func (fortune) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && (t.ToolType() == item.TypePickaxe || t.ToolType() == item.TypeShovel || t.ToolType() == item.TypeAxe || t.ToolType() == item.TypeHoe)
}
