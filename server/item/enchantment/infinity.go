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

// Infinity is an enchantment to bows that prevents regular arrows from being
// consumed when shot.
var Infinity infinity

type infinity struct{}

// Name ...
func (infinity) Name() string {
	return "Infinity"
}

// MaxLevel ...
func (infinity) MaxLevel() int {
	return 1
}

// Cost ...
func (infinity) Cost(int) (int, int) {
	return 20, 50
}

// Rarity ...
func (infinity) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// ConsumesArrows always returns false.
func (infinity) ConsumesArrows() bool {
	return false
}

// CompatibleWithEnchantment ...
func (infinity) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != Mending
}

// CompatibleWithItem ...
func (infinity) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Bow)
	return ok
}
