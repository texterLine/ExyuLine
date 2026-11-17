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

// Respiration extends underwater breathing time by +15 seconds per enchantment
// level in addition to the default time of 15 seconds.
var Respiration respiration

type respiration struct{}

// Name ...
func (respiration) Name() string {
	return "Respiration"
}

// MaxLevel ...
func (respiration) MaxLevel() int {
	return 3
}

// Cost ...
func (respiration) Cost(level int) (int, int) {
	minCost := 10 * level
	return minCost, minCost + 30
}

// Rarity ...
func (respiration) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// Chance returns the chance of the enchantment blocking the air supply from ticking.
func (respiration) Chance(level int) float64 {
	return float64(level) / float64(level+1)
}

// CompatibleWithEnchantment ...
func (respiration) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (respiration) CompatibleWithItem(i world.Item) bool {
	h, ok := i.(item.HelmetType)
	return ok && h.Helmet()
}
