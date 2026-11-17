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

// AquaAffinity is a helmet enchantment that increases underwater mining speed.
var AquaAffinity aquaAffinity

type aquaAffinity struct{}

// Name ...
func (aquaAffinity) Name() string {
	return "Aqua Affinity"
}

// MaxLevel ...
func (aquaAffinity) MaxLevel() int {
	return 1
}

// Cost ...
func (aquaAffinity) Cost(int) (int, int) {
	return 1, 41
}

// Rarity ...
func (aquaAffinity) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (aquaAffinity) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (aquaAffinity) CompatibleWithItem(i world.Item) bool {
	h, ok := i.(item.HelmetType)
	return ok && h.Helmet()
}
