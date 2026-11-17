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

// Multishot is an enchantment for crossbows that allow them to shoot three arrows or firework rockets at the cost of one.
var Multishot multishot

type multishot struct{}

// Name ...
func (multishot) Name() string {
	return "Multishot"
}

// MaxLevel ...
func (multishot) MaxLevel() int {
	return 1
}

// Cost ...
func (m multishot) Cost(level int) (int, int) {
	return 20, 50
}

// Rarity ...
func (multishot) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (multishot) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != Piercing
}

// CompatibleWithItem ...
func (multishot) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Crossbow)
	return ok
}

// MultipleProjectiles ...
func (multishot) MultipleProjectiles() bool {
	return true
}
