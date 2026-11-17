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

// Flame turns your arrows into flaming arrows allowing you to set your targets
// on fire.
var Flame flame

type flame struct{}

// Name ...
func (flame) Name() string {
	return "Flame"
}

// MaxLevel ...
func (flame) MaxLevel() int {
	return 1
}

// Cost ...
func (flame) Cost(int) (int, int) {
	return 20, 50
}

// Rarity ...
func (flame) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// BurnDuration always returns five seconds, no matter the level.
func (flame) BurnDuration() time.Duration {
	return time.Second * 5
}

// CompatibleWithEnchantment ...
func (flame) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (flame) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Bow)
	return ok
}
