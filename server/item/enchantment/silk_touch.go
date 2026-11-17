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

// SilkTouch is an enchantment that allows many blocks to drop themselves
// instead of their usual items when mined.
var SilkTouch silkTouch

type silkTouch struct{}

// Name ...
func (silkTouch) Name() string {
	return "Silk Touch"
}

// MaxLevel ...
func (silkTouch) MaxLevel() int {
	return 1
}

// Cost ...
func (silkTouch) Cost(int) (int, int) {
	return 15, 65
}

// Rarity ...
func (silkTouch) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// CompatibleWithEnchantment ...
func (silkTouch) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != Fortune
}

// CompatibleWithItem ...
func (silkTouch) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && (t.ToolType() != item.TypeSword && t.ToolType() != item.TypeNone)
}
