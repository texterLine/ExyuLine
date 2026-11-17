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

// CurseOfVanishing is an enchantment that causes the item to disappear on
// death.
var CurseOfVanishing curseOfVanishing

type curseOfVanishing struct{}

// Name ...
func (curseOfVanishing) Name() string {
	return "Curse of Vanishing"
}

// MaxLevel ...
func (curseOfVanishing) MaxLevel() int {
	return 1
}

// Cost ...
func (curseOfVanishing) Cost(int) (int, int) {
	return 25, 50
}

// Rarity ...
func (curseOfVanishing) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// CompatibleWithEnchantment ...
func (curseOfVanishing) CompatibleWithEnchantment(_ item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (curseOfVanishing) CompatibleWithItem(i world.Item) bool {
	_, arm := i.(item.Armour)
	_, com := i.(item.Compass)
	_, dur := i.(item.Durable)
	_, rec := i.(item.RecoveryCompass)
	// TODO: Carrot on a Stick
	// TODO: Warped Fungus on a Stick
	return arm || com || dur || rec
}

// Treasure ...
func (curseOfVanishing) Treasure() bool {
	return true
}

// Curse ...
func (curseOfVanishing) Curse() bool {
	return true
}
