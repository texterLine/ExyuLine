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

// GoldenApple is a special food item that bestows beneficial effects.
type GoldenApple struct{}

// AlwaysConsumable ...
func (e GoldenApple) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (e GoldenApple) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (e GoldenApple) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(4, 9.6)
	prev := c.Absorption()
	c.AddEffect(effect.New(effect.Absorption, 1, 2*time.Minute))
	c.SetAbsorption(max(prev, min(prev+4, 16)))
	c.AddEffect(effect.New(effect.Regeneration, 2, 5*time.Second))
	return Stack{}
}

// EncodeItem ...
func (e GoldenApple) EncodeItem() (name string, meta int16) {
	return "minecraft:golden_apple", 0
}
