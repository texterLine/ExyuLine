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

// FireAspect is a sword enchantment that sets the target on fire.
var FireAspect fireAspect

type fireAspect struct{}

// Name ...
func (fireAspect) Name() string {
	return "Fire Aspect"
}

// MaxLevel ...
func (fireAspect) MaxLevel() int {
	return 2
}

// Cost ...
func (fireAspect) Cost(level int) (int, int) {
	minCost := 10 + (level-1)*20
	return minCost, minCost + 50
}

// Rarity ...
func (fireAspect) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// Duration returns how long the fire from fire aspect will last.
func (fireAspect) Duration(level int) time.Duration {
	return time.Second * 4 * time.Duration(level)
}

// CompatibleWithEnchantment ...
func (fireAspect) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (fireAspect) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && t.ToolType() == item.TypeSword
}
