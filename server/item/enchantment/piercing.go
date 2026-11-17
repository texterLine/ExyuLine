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

// Piercing is an enchantment that allows arrows to damage and pierce through multiple entities, including shields.
var Piercing piercing

type piercing struct{}

func (p piercing) Name() string {
	return "Piercing"
}

func (p piercing) MaxLevel() int {
	return 4
}

func (p piercing) Cost(level int) (int, int) {
	return 1 + (level-1)*10, 50
}

func (p piercing) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityCommon
}

func (p piercing) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return t != Multishot
}

func (p piercing) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Crossbow)
	return ok
}

func (p piercing) Pierces() bool {
	return true
}
