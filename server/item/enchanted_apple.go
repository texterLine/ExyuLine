// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/entity/effect"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// EnchantedApple is a rare variant of the golden apple that has stronger effects.
type EnchantedApple struct{}

// AlwaysConsumable ...
func (EnchantedApple) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (EnchantedApple) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (EnchantedApple) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(4, 9.6)
	c.AddEffect(effect.New(effect.Absorption, 4, 2*time.Minute))
	c.AddEffect(effect.New(effect.Regeneration, 2, 30*time.Second))
	c.AddEffect(effect.New(effect.FireResistance, 1, 5*time.Minute))
	c.AddEffect(effect.New(effect.Resistance, 1, 5*time.Minute))
	return Stack{}
}

// EncodeItem ...
func (EnchantedApple) EncodeItem() (name string, meta int16) {
	return "minecraft:enchanted_golden_apple", 0
}
