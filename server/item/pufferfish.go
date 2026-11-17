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

// Pufferfish is a poisonous type of fish that is used to brew water breathing potions.
type Pufferfish struct {
	defaultFood
}

// Consume ...
func (p Pufferfish) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(1, 0.2)
	c.AddEffect(effect.New(effect.Hunger, 3, 15*time.Second))
	c.AddEffect(effect.New(effect.Poison, 2, time.Minute))
	c.AddEffect(effect.New(effect.Nausea, 2, 15*time.Second))
	return Stack{}
}

// EncodeItem ...
func (p Pufferfish) EncodeItem() (name string, meta int16) {
	return "minecraft:pufferfish", 0
}
