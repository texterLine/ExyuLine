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
	"time"
)

// QuickCharge is an enchantment for quickly reloading a crossbow.
var QuickCharge quickCharge

type quickCharge struct{}

// Name ...
func (quickCharge) Name() string {
	return "Quick Charge"
}

// MaxLevel ...
func (quickCharge) MaxLevel() int {
	return 3
}

// Cost ...
func (quickCharge) Cost(level int) (int, int) {
	minCost := 12 + (level-1)*20
	return minCost, 50
}

// Rarity ...
func (quickCharge) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// ChargeDuration returns the charge duration.
func (quickCharge) ChargeDuration(level int) time.Duration {
	return time.Duration((1.25 - 0.25*float64(level)) * float64(time.Second))
}

// CompatibleWithEnchantment ...
func (quickCharge) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (quickCharge) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Crossbow)
	return ok
}
