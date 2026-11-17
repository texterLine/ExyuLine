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
	"math/rand/v2"
	"time"
)

// RottenFlesh is a food item that can be eaten by the player, at the high risk of inflicting hunger.
type RottenFlesh struct {
	defaultFood
}

// Consume ...
func (RottenFlesh) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(4, 0.8)
	if rand.Float64() < 0.8 {
		c.AddEffect(effect.New(effect.Hunger, 1, 30*time.Second))
	}
	return Stack{}
}

// EncodeItem ...
func (RottenFlesh) EncodeItem() (name string, meta int16) {
	return "minecraft:rotten_flesh", 0
}
