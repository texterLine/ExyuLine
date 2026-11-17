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

// Sharpness is an enchantment applied to a sword or axe that increases melee
// damage.
var Sharpness sharpness

type sharpness struct{}

// Name ...
func (sharpness) Name() string {
	return "Sharpness"
}

// MaxLevel ...
func (sharpness) MaxLevel() int {
	return 5
}

// Cost ...
func (sharpness) Cost(level int) (int, int) {
	minCost := 1 + (level-1)*11
	return minCost, minCost + 20
}

// Rarity ...
func (sharpness) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityCommon
}

// Addend returns the additional damage when attacking with sharpness.
func (sharpness) Addend(level int) float64 {
	return float64(level) * 1.25
}

// CompatibleWithEnchantment ...
func (sharpness) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (sharpness) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && (t.ToolType() == item.TypeSword || t.ToolType() == item.TypeAxe)
}
